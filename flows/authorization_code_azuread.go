package flows

import (
	"errors"

	"github.com/izassan/oauth2-testtool/types"
)

type azureADAuthorizer struct{}

func (a *azureADAuthorizer) init(config *types.OttConfig) error{
    return errors.New("Unimplemented")
}

func (a *azureADAuthorizer) GenerateAuthorizeURL(config *types.OttConfig) (string, error){
    return "", errors.New("Unimplemented")
}

func (a *azureADAuthorizer) ExchangeAccessToken(authorizationCode string) (*token, error){
    return nil, errors.New("Unimplemented")
}

func (a *azureADAuthorizer) ValidateState(state string) error{
    return errors.New("Unimplemented")
}
