package code

import (
	"context"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type idToken struct{
    jwt.Token
}

func parseIdToken(jwtString string, jwkURI string, correctNonce string) (*idToken, error) {
    keyset, err := getJWKKey(jwkURI)
    if err != nil{
        return nil, err
    }

    token, err := jwt.Parse([]byte(jwtString), jwt.WithKeySet(keyset))
    if err != nil{
        return nil, err
    }
    return &idToken{token}, nil
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
