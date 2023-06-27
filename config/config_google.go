package config

import (
	"github.com/izassan/oidc-testing-tool/security"
)

const GOOGLE_PROVIDER_NAME = "google"

type GoogleOTTConfig struct{
    AuthFlow string `json:"auth_flow"`
    ClientId string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
    Scope string `json:"scope"`
    Authorize GoogleAuthorizeParams `json:"authorize"`
    Token GoogleTokenParams `json:"token"`
    JwkURI string `json:"jwk_uri"`
    SecurityParams *security.SecurityParams
}

type GoogleAuthorizeParams struct{
    AuthorizeURI string `json:"auth_uri"`
    AuthorizeOptions GoogleAuthorizeOptions `json:"options"`
}

type GoogleAuthorizeOptions struct{
    AccessType string `json:"access_type"`
    Display string `json:"display"`
    Hd string `json:"hd"`
    LoginHint string `json:"login_hint"`
    Prompt string `json:"prompt"`
}

type GoogleTokenParams struct{
    TokenURI string `json:"token_uri"`
}
