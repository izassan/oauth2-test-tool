package config

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/izassan/oidc-testing-tool/security"
)

const GOOGLE_PROVIDER_NAME = "google"

type GoogleOTTConfig struct{
    AuthFlow string `json:"auth_flow"`
    ClientId string `json:"client_id" qname:"client_id"`
    ClientSecret string `json:"client_secret" qname:"client_secret"`
    Scope string `json:"scope" qname:"scope"`
    Authorize AuthorizeParams `json:"authorize"`
    Token TokenParams `json:"token"`
    JwkURI string `json:"jwk_uri"`
    SecurityParams *security.SecurityParams
}

type AuthorizeParams struct{
    AuthorizeURI string `json:"auth_uri"`
    AuthorizeOptions AuthorizeOptions `json:"options"`
}

type AuthorizeOptions struct{
    AccessType string `json:"access_type" qname:"access_type"`
    Display string `json:"display" qname:"display"`
    Hd string `json:"hd" qname:"hd"`
    LoginHint string `json:"login_hint" qname:"login_hint"`
    Prompt string `json:"prompt" qname:"prompt"`
}

type TokenParams struct{
    TokenURI string `json:"token_uri"`
}

func (c GoogleOTTConfig) GetFullAuthorizeURI() (string, error){
    authorizeUri, err := url.Parse(c.Authorize.AuthorizeURI)
    if err != nil{
        return "", err
    }
    q := authorizeUri.Query()
    // TODO: add queries
    // required: client_id, redirect_uri, scope, response_type
    // optional: state, nonce, code_challenge, code_challenge_method,
    //              access_type, display, hd, login_hint, prompt

    authorizeUri.RawQuery = q.Encode()
    return authorizeUri.String(), nil

}

func (c GoogleOTTConfig) GetTokenRequest() (*http.Request, error){
    q := url.Values{}
    // TODO: add queries
    // required: code, client_id, client_secret, grant_type
    // optional: code_verifier

    req, err := http.NewRequest("POST", c.Token.TokenURI, bytes.NewReader([]byte(q.Encode())))
    if err != nil{
        return nil, err
    }
    return req, nil
}

func (c GoogleOTTConfig) CreateSecurityParams() error{
    sp, err := security.NewSecurityParams()
    if err != nil{
        return err
    }
    c.SecurityParams = sp
    return nil
}
