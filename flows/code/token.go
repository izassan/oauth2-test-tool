package code

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"context"
    "fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type authorizeParameters struct {
    clientId string
    scope string
    responseType string
    redirectURI string
    state string
    nonce string
    pkce *pkce
}

type authorizeCodeParameters struct {
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
type IDToken struct{
    jwt.Token
}

type Token struct {
    AccessToken string `json:"access_token"`
    ExpiresIn int `json:"expires_in"`
    IdToken string `json:"id_token"`
    Scope string `json:"scope"`
    TokenType string `json:"token_type"`
    RefreshToken string `json:"refresh_token"`
}

func parseIdToken(jwtString string, jwkURI string, correctNonce string) (*IDToken, error) {
    keyset, err := getJWKKey(jwkURI)
    if err != nil{
        return nil, err
    }

    token, err := jwt.Parse([]byte(jwtString), jwt.WithKeySet(keyset))
    if err != nil{
        return nil, err
    }
    return &IDToken{token}, nil
}

func getJWKKey(jwkURI string) (jwk.Set, error) {
    keyset, err := jwk.Fetch(context.Background(), jwkURI)
    if err != nil{
        return nil, err
    }

    for it := keyset.Iterate(context.Background()); it.Next(context.Background()); {
        pair := it.Pair()
        key := pair.Value.(jwk.Key)

        var rawkey interface{}
        if err := key.Raw(&rawkey); err != nil{
            return nil, err
        }
    }
    return keyset, nil
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

func generateAuthorizeURL(authURL string, authParams *authorizeParameters) (string, error){
    authorizeURL, err := url.Parse(authURL)
    if err != nil{
        return "", err
    }

    q := authorizeURL.Query()
    q.Add("client_id", authParams.clientId)
    q.Add("redirect_uri", authParams.redirectURI)
    q.Add("scope", authParams.scope)
    q.Add("response_type", authParams.responseType)

    if authParams.state != ""{
        q.Add("state", authParams.state)
    }
    if authParams.nonce != ""{
        q.Add("nonce", authParams.nonce)
    }
    if authParams.pkce != nil{
        q.Add("code_challenge", authParams.pkce.codeChallenge)
        q.Add("code_challenge_method", authParams.pkce.codeChallengeMethod)
    }

    authorizeURL.RawQuery = q.Encode()

    return authorizeURL.String(), nil
}

func outputResult(t *Token, it *IDToken, formatType string){
    if formatType != "default"{
        fmt.Println("warning: specified unsupported or unimplement output format. output default format")
    }
    t.outputToken()
    if it != nil{
        fmt.Printf("----------------------\n")
        it.outputIDToken()
    }
}

func (t *Token)outputToken() {
	fmt.Printf("access_token: %s\n", t.AccessToken)
	fmt.Printf("id_token: %s\n", t.IdToken)
	fmt.Printf("refresh_token: %s\n", t.RefreshToken)
	fmt.Printf("scope: %s\n", t.Scope)
	fmt.Printf("token_type: %s\n", t.TokenType)
	fmt.Printf("expire_in: %d\n", t.ExpiresIn)
}

func (t *IDToken)outputIDToken(){
        fmt.Printf("id_token audiences: %s\n", t.Audience())
        fmt.Printf("id_token issuer: %s\n", t.Issuer())
}
