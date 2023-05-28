package flows

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/izassan/oauth2-testtool/types"
	"github.com/spf13/pflag"
)

type authorizer interface {
    init(config *types.OttConfig) error
	GenerateAuthorizeURL(config *types.OttConfig) (string, error)
	ExchangeAccessToken(authorizationCode string) (*token, error)
	ValidateState(state string) error
}

type callbackQueries struct {
    authorizeCode string
    state string
    scope string
}

type token struct {
	accessToken     string
	refreshToken    string
	expireIn        time.Time
	expireInRefresh time.Time
}

func AuthorizationCodeFlow(config *types.OttConfig, flags *pflag.FlagSet) error {
    useAuthorizer, err := selectAuthorizer(config.Provider)
    if err != nil{
        return err
    }
    if err := useAuthorizer.init(config); err != nil{
        return err
    }

    authorizeEndpointURL, err := useAuthorizer.GenerateAuthorizeURL(config)
    if err != nil{
        return err
    }
    if config.Require_string == "authorize_url"{
        fmt.Printf("Authorize Endpoint URL: %s\n", authorizeEndpointURL)
        return nil
    }

	noBrowser, err := flags.GetBool("no-browser")
	if err != nil {
		return err
	}

	served := make(chan bool)
	if noBrowser {
		go outputAuthorizeURL(authorizeEndpointURL, served)
	} else {
		go openBrowser(authorizeEndpointURL, served)
	}

	ccq := make(chan *callbackQueries)
	port, err := flags.GetInt("port")
	go startCallbackServer(port, ccq, served)
    cq := <-ccq
    if config.Require_string == "authorize_code"{
        fmt.Printf("Authorize Endpoint Response:\n")
        fmt.Printf("\tcode: \"%s\"\n", cq.authorizeCode)
        fmt.Printf("\tscope: \"%s\"\n", cq.scope)
        fmt.Printf("\tstate: \"%s\"\n", cq.state)
        return nil
    }

    if err := useAuthorizer.ValidateState(cq.state); err != nil{
        return err
    }

    token, err := useAuthorizer.ExchangeAccessToken(cq.authorizeCode)
    if err != nil{
        return err
    }

    if config.Require_string == "token"{
        fmt.Printf("Token Endpoint Response:\n")
        fmt.Printf("\taccess_token: \"%s\"\n", token.accessToken)
        fmt.Printf("\trefresh_token: \"%s\"\n", token.refreshToken)
        fmt.Printf("\texpire: \"%s\"\n", token.expireIn.Format("2006-01-02 15:04"))
    }
	return nil
}

func selectAuthorizer(provider string) (authorizer, error) {
    if provider == "google"{
        return &googleAuthorizer{}, nil
    }
	return nil, errors.New("error")
}

func outputAuthorizeURL(authorizeURL string, served chan bool) {
	<-served
	fmt.Printf("access to: \n\t%s\n\n", authorizeURL)
}

func openBrowser(authorizeURL string, served chan bool) {
	<-served
	fmt.Printf("open authorize endpoint with browser...\n\n")
}

func startCallbackServer(port int, cq chan *callbackQueries, served chan bool) error {
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{Addr: addr}
	sig := make(chan bool)

	handler := func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello world\n")
        rcq := &callbackQueries{
            authorizeCode: r.URL.Query().Get("code"),
            state: r.URL.Query().Get("state"),
            scope: r.URL.Query().Get("scope"),
        }
		cq <- rcq
		sig <- true
	}
	http.HandleFunc("/callback", handler)

	go func() {
		served <- true
		if err := server.ListenAndServe(); err != nil {
		}
	}()
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
