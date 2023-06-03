package code

import (
	"fmt"

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
    authParam := &authorizeParameters{
        clientId: config.ClientId,
        scope: config.Scope,
        responseType: RESPONSETYPE,
        redirectURI: fmt.Sprintf("http://%s:%d/callback", host, port),
        state: generateState(),
        nonce: generateNonce(),
        pkce: generatePKCE(),
    }

    // TODO: generate AuthorizeURL
    authURI, err := generateAuthorizeURL(authParam)

    // TODO: output stdout or start browser
    fmt.Printf(authURI)

    // TODO: start callback server
    go startCallbackServer()

    // TODO: get request to redirect_uri

    // TODO: request to token endpoint

    // TODO: output result
    return nil
}