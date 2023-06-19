package code

import (
	"context"
    "fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Token struct {
    AccessToken string `json:"access_token"`
    ExpiresIn int `json:"expires_in"`
    IdToken string `json:"id_token"`
    Scope string `json:"scope"`
    TokenType string `json:"token_type"`
    RefreshToken string `json:"refresh_token"`
    DecodedIdToken *IDToken `json:"id_token_decoded"`
}

type IDToken struct{
    jwt.Token
}

func (t *Token)ParseIdToken(jwtString string, jwkURI string, correctNonce string) (error) {
    keyset, err := jwk.Fetch(context.Background(), jwkURI)
    if err != nil{
        return err
    }

    token, err := jwt.Parse([]byte(jwtString), jwt.WithKeySet(keyset))
    if err != nil{
        return err
    }
    t.DecodedIdToken = &IDToken{token}
    return nil
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
