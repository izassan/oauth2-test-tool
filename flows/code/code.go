package code

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type AuthorizeCode struct {
    code string
    state string
    scope string
}

type (
    VerifyParams struct{
        state string
    }
    AuthorizeValifyOption func(*VerifyParams)error
)
func VerifyState(correctState string) AuthorizeValifyOption{
    return func(v *VerifyParams)error{
        v.state = correctState
        return nil
    }
}

type tokenExchangeParams struct {
    tokenURI string
    clientId string
    clientSecret string
    redirectURI string
    grantType string
    codeVerifier string
}


func NewAuthorizeCode(r *http.Request, options ...AuthorizeValifyOption) (*AuthorizeCode, error){
    config := &VerifyParams{
        state: "",
    }

    for _, opt := range options{
        opt(config)
    }

    queries := r.URL.Query()
    authCode := &AuthorizeCode{
        code: queries.Get("code"),
        state: queries.Get("state"),
        scope: queries.Get("scope"),
    }


    if config.state != "" && config.state != authCode.state{
        return nil, errors.New("invalid state parameter")
    }

    return authCode, nil
}

func (c *AuthorizeCode)ExchangeToken(params *tokenExchangeParams) (*Token, error){
    q := url.Values{}
    q.Add("code", c.code)
    q.Add("client_id", params.clientId)
    q.Add("client_secret", params.clientSecret)
    q.Add("redirect_uri", params.redirectURI)
    q.Add("grant_type", params.grantType)
    q.Add("code_verifier", params.codeVerifier)

    res, err := http.PostForm(params.tokenURI, q)
    if err != nil{
        return nil, err
    }
    defer res.Body.Close()

    bodyBytes, _ := io.ReadAll(res.Body)
    var t Token
    if err := json.Unmarshal(bodyBytes, &t); err != nil{
        return nil, err
    }
    return &t, nil
}
