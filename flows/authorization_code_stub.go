package flows

import (
    "fmt"
	"net/http"

	"github.com/izassan/oauth2-testtool/types"
)

func AuthorizationCodeFlow(config *types.OttConfig, no_browser bool) error{
	AuthorizeRequest, err := generateAuthorizeURL(config)
	if err != nil {
		return err
	}

	c := make(chan string)
    served := make(chan bool)
    if no_browser{
        go outputAuthorizeURL(AuthorizeRequest, served)
    }else{
        go openAuthorize(AuthorizeRequest, served)
    }
	go startCallbackServer(c, served)
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

func startCallbackServer(c chan string, served chan bool){
    fmt.Println("startCallbackServer")
    served <- true
    c <- "authorize_code"
}

func exhangeAccessToken(authorizeCode string, config *types.OttConfig) (string, error){
    fmt.Println("exhangeAccessToke")
	return "access_token", nil
}
