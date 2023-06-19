package code
import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type AuthorizeCode struct {
    code string
    state string
    scope string
}

type tokenExchangeParams struct {
    code string
    clientId string
    clientSecret string
    redirectURI string
    grantType string
    codeVerifier string
}

func exchangeToken(tokenURI string, params *tokenExchangeParams) (*Token, error){
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
    var t Token
    if err := json.Unmarshal(bodyBytes, &t); err != nil{
        return nil, err
    }
    return &t, nil
}
