package code

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/izassan/oidc-testing-tool/config"
	"github.com/spf13/pflag"
)

const RESPONSETYPE = "code"

func ExecuteAuthorizeCodeFlow(config *config.OttConfig, flags *pflag.FlagSet) error{
    host, err := flags.GetString("host")
    if err != nil{
        return err
    }
    port, err := flags.GetInt("port")
    if err != nil{
        return err
    }

    state, err := generateState()
    if err != nil{
        return err
    }

    nonce, err := generateNonce()
    if err != nil{
        return err
    }

    pkce, err := generatePKCE()
    if err != nil{
        return err
    }

    authParam := &authorizeParameters{
        clientId: config.ClientId,
        scope: config.Scope,
        responseType: RESPONSETYPE,
        redirectURI: fmt.Sprintf("http://%s:%d/callback", host, port),
        state: state,
        nonce: nonce,
        pkce: pkce,
    }

    authURI, err := generateAuthorizeURL(config.AuthURI, authParam)

    noUseBrowser, err := flags.GetBool("no-browser")
    if err != nil{
        return err
    }
    if !noUseBrowser{
        browserError := make(chan error)
        go startBrowser(authURI, browserError)
        be := <-browserError
        if be != nil{
            return err
        }
    }else{
        fmt.Printf("Access To:\n\t%s\n", authURI)
    }


    ctx, cancel := context.WithCancel(context.Background())
    callbackRequestChannel := make(chan *http.Request)
    go startCallbackServer(ctx, host, port, callbackRequestChannel)
    callbackRequest := <- callbackRequestChannel
    cancel()

    queries := callbackRequest.URL.Query()
    authCodeParams := &authorizeCodeParameters{
        code: queries.Get("code"),
        state: queries.Get("state"),
        scope: queries.Get("scope"),
    }
    if authCodeParams.state != state{
        return errors.New("invlid state")
    }

    exchangeParams := &tokenExchangeParams{
        code: authCodeParams.code,
        clientId: config.ClientId,
        clientSecret: config.ClientSecret,
        redirectURI: authParam.redirectURI,
        grantType: "authorization_code",
        codeVerifier: authParam.pkce.codeVerifier,
    }
    token, err := exchangeToken(config.TokenURI, exchangeParams)
    if err != nil{
        return err
    }

    decodedIdToken, err := parseIdToken(token.IdToken, nonce)
    if err != nil{
        return err
    }

    fmt.Printf("access_token: %s\n", token.AccessToken)
    fmt.Printf("id_token_raw: %s\n", token.IdToken)
    fmt.Printf("id_token_nonce: %s\n", decodedIdToken.claims["nonce"])
    fmt.Printf("refresh_token: %s\n", token.RefreshToken)
    fmt.Printf("scope: %s\n", token.Scope)
    fmt.Printf("token_type: %s\n", token.TokenType)
    fmt.Printf("expire_in: %d\n", token.ExpiresIn)

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
