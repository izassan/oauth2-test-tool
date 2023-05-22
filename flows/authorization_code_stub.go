package flows

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/izassan/oauth2-testtool/types"
	"github.com/spf13/pflag"
)

func AuthorizationCodeFlow(config *types.OttConfig, flags *pflag.FlagSet) error{
	AuthorizeRequest, err := generateAuthorizeURL(config)
	if err != nil {
		return err
	}

	c := make(chan string)
    served := make(chan bool)
    no_browser, err := flags.GetBool("no-browser")
    if err != nil{
        return err
    }

    if no_browser{
        go outputAuthorizeURL(AuthorizeRequest, served)
    }else{
        go openAuthorize(AuthorizeRequest, served)
    }

    port, err := flags.GetInt("port")
    if err != nil{
        return err
    }
	go startCallbackServer(port, c, served)
	authorizeCode := <-c

	token, err := exhangeAccessToken(authorizeCode, config)
	fmt.Println(token)
    return nil
}

func generateAuthorizeURL(config *types.OttConfig) (*http.Request, error) {
    fmt.Println("generateAuthorizeURL")
	return nil, nil
}

func outputAuthorizeURL(request *http.Request, served chan bool){
    <-served
    fmt.Println("outputAuthorizeURL")
}

func openAuthorize(request *http.Request, served chan bool){
    <-served
    fmt.Println("openAuthorize")
}

func startCallbackServer(port int, c chan string, served chan bool) error {
    addr := fmt.Sprintf(":%d", port)
    server := &http.Server{Addr: addr}
    sig := make(chan bool)

    handler := func(w http.ResponseWriter, r *http.Request){
        fmt.Println("hello world")
        c <- "authorize_code"
        sig <- true
    }
    http.HandleFunc("/callback", handler)

    go func(){
        served <- true
        if err := server.ListenAndServe(); err != nil{
            log.Print(err)
        }
    }()
    <-sig


    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil{
        return err
    }
    return nil
}

func exhangeAccessToken(authorizeCode string, config *types.OttConfig) (string, error){
    fmt.Println("exhangeAccessToke")
	return "access_token", nil
}
