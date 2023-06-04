package code

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func exchangeToken(tokenURI string, params *tokenExchangeParams) (*token, error){
    q := url.Values{}
    q.Add("code", params.code)
    q.Add("client_id", params.clientId)
    q.Add("client_secret", params.clientSecret)
    q.Add("redirect_uri", params.redirectURI)
    q.Add("grant_type", params.grantType)
    q.Add("code_verifier", params.codeVerifier)

    res, err := http.PostForm(tokenURI, q)
    if err != nil{
        return nil, err
    }
    defer res.Body.Close()

    bodyBytes, _ := io.ReadAll(res.Body)
    body := string(bodyBytes)
    fmt.Println(body)
    return nil, errors.New("unimpletement")
}
