package flows

import (
	"context"
	"errors"
	"os"

	"github.com/izassan/oauth2-testtool/types"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleAuthorizer struct{
    googleConfig *oauth2.Config
    state string
}

func (g *googleAuthorizer) init(config *types.OttConfig) error {
    CredentialFile, err := os.Open(config.CredentialFile)
    if err != nil{
        return err
    }
    CredentialData := make([]byte, 1024)
    count, err := CredentialFile.Read(CredentialData)
    if err != nil{
        return err
    }
    credentialBytes := CredentialData[:count]

    googleConfig, err := google.ConfigFromJSON(credentialBytes, config.Scope)
    googleConfig.RedirectURL = "http://localhost:8893/callback"
    if err != nil{
        return err
    }
    g.googleConfig = googleConfig
    return nil
}

func (g *googleAuthorizer) GenerateAuthorizeURL(config *types.OttConfig) (string, error) {
    g.state = "oauth2-test-tool-google-authorization-flow-state"
    return g.googleConfig.AuthCodeURL(g.state, oauth2.AccessTypeOffline, oauth2.ApprovalForce), nil
}

func (g *googleAuthorizer) ExchangeAccessToken(authorizationCode string) (*token, error){
    ctx := context.Background()
    t, err := g.googleConfig.Exchange(ctx, authorizationCode)
    if err != nil{
        return nil, err
    }

    return &token{
        accessToken: t.AccessToken,
        refreshToken: t.RefreshToken,
        expireIn: t.Expiry,
    }, nil
}

func (g *googleAuthorizer) ValidateState(state string) error{
    if state == g.state{
        return nil
    }
    return errors.New("invalid state parameter")
}
