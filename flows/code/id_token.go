package code

import (
	"context"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type idToken struct{
    jwt.Token
}

func parseIdToken(jwtString string, correctNonce string) (*idToken, error) {
    const jwk_uri="https://www.googleapis.com/oauth2/v3/certs"
    keyset, err := getJWKKey(jwk_uri)
    if err != nil{
        return nil, err
    }

    token, err := jwt.Parse([]byte(jwtString), jwt.WithKeySet(keyset))
    if err != nil{
        return nil, err
    }
    return &idToken{token}, nil
}

func getJWKKey(jwk_uri string) (jwk.Set, error) {
    keyset, err := jwk.Fetch(context.Background(), jwk_uri)
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
