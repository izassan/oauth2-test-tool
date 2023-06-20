package code

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

)

const RESPONSETYPE = "code"

func ExecuteAuthorizeCodeFlow(config *AuthorizeCodeFlowConfig) error{
    sp, err := newSecurityParams()
    if err != nil{
        return err
    }

    authParam := &authorizeParameters{
        clientId: config.ClientId,
        scope: config.Scope,
        responseType: RESPONSETYPE,
        redirectURI: fmt.Sprintf("http://%s:%d/callback", config.RPConfig.Host, config.RPConfig.Port),
        state: sp.state,
        nonce: sp.nonce,
        pkce: sp.pkce,
    }

    authURI, err := generateAuthorizeURL(config.AuthURI, authParam)

    if config.UseBrowser{
        browserError := make(chan error)
        go startBrowser(authURI, browserError)
        be := <-browserError
        if be != nil{
            return err
        }
    }else{
        fmt.Printf("Access To:\n\t%s\n\n", authURI)
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    callbackRequestChannel := make(chan *http.Request)
    go startCallbackServer(ctx, config.RPConfig.Host, config.RPConfig.Port, callbackRequestChannel)
    callbackRequest := <- callbackRequestChannel

    authorizeCode, err  := NewAuthorizeCode(callbackRequest, VerifyState(sp.state))
    if err != nil{
        return err
    }

    exchangeParams := &tokenExchangeParams{
        tokenURI: config.TokenURI,
        clientId: config.ClientId,
        clientSecret: config.ClientSecret,
        redirectURI: authParam.redirectURI,
        grantType: "authorization_code",
        codeVerifier: authParam.pkce.codeVerifier,
    }
    token, err := authorizeCode.ExchangeToken(exchangeParams)
    if err != nil{
        return err
    }

    if config.RequiredVerify{
        err := token.ParseIdToken(token.IdToken, config.JwkURI, sp.nonce)
        if err != nil{
            return err
        }
        outputResult(token, token.DecodedIdToken, "default")
    }else{
        outputResult(token, nil, "default")
    }
    return nil
}

func startBrowser(uri string, berr chan error){
    var err error
	switch runtime.GOOS {
	case "linux":
        err = exec.Command("xdg-open", uri).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", uri).Start()
	case "darwin":
		err = exec.Command("open", uri).Start()
	default:
		err = errors.New("Unsupported OS")
	}
    if err != nil{
        berr <- err
        return
    }
    berr <- nil
}

func outputResult(t *Token, it *IDToken, formatType string){
    if formatType != "default"{
        fmt.Println("warning: specified unsupported or unimplement output format. output default format")
    }
    t.OutputToken()
    if it != nil{
        fmt.Printf("----------------------\n")
        it.OutputIDToken()
    }
}
