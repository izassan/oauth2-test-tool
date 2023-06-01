package flows

import (
	"context"
	"errors"
	"strings"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/izassan/oauth2-testtool/types"
)

type azureADAuthorizer struct{
    clientApp *public.Client
    ctx context.Context
}

func (a *azureADAuthorizer) init(config *types.OttConfig) error{
    publicApp, err := public.New(config.Client_id, public.WithAuthority(""))
    if err != nil{
        return err
    }
    a.clientApp = &publicApp
    a.ctx = context.Background()
    return errors.New("Unimplemented")
}

func (a *azureADAuthorizer) GenerateAuthorizeURL(config *types.OttConfig, redirectURL string) (string, error){
    scopes := strings.Split(config.Scope, " ")
    authorizeURL, err := a.clientApp.AuthCodeURL(a.ctx, config.Client_id, redirectURL, scopes)
    if err != nil{
        return "", err
    }
    return authorizeURL, nil
}

func (a *azureADAuthorizer) ExchangeAccessToken(authorizationCode string) (*token, error){
    return nil, errors.New("Unimplemented")
}

func (a *azureADAuthorizer) ValidateState(state string) error{
    return errors.New("Unimplemented")
}
