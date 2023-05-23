package flows

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/izassan/oauth2-testtool/types"
	"github.com/spf13/pflag"
)

func AuthorizationCodeFlow(config *types.OttConfig, flags *pflag.FlagSet) error{
    // TODO: create authorize endpoint URL

    noBrowser, err := flags.GetBool("no-browser")
    if err != nil{
        return err
    }

    served := make(chan bool)
    if noBrowser{
        go outputAuthorizeURL("authorizeURL", served)
    }else{
        go openBrowser("authorizeURL", served)
    }

    c := make(chan string)
    port, err := flags.GetInt("port")
    go startCallbackServer(port, c, served)
    fmt.Printf("authorize_code: %s", <-c)

    // TODO: exchange access token
    return nil
}

func outputAuthorizeURL(authorizeURL string, served chan bool){
    <-served
    fmt.Printf("access to: \n\t%s\n\n", authorizeURL)
}

func openBrowser(authorizeURL string, served chan bool){
    <-served
    fmt.Printf("open authorize endpoint with browser...\n\n")
}

func startCallbackServer(port int, c chan string, served chan bool) error {
    addr := fmt.Sprintf(":%d", port)
    server := &http.Server{Addr: addr}
    sig := make(chan bool)

    handler := func(w http.ResponseWriter, r *http.Request){
        c <- "authorize_code"
        sig <- true
    }
    http.HandleFunc("/callback", handler)

    go func(){
        served <- true
        if err := server.ListenAndServe(); err != nil{
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
