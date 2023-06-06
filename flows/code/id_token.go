package code

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type idToken struct{
    header map[string]string
    claims jwt.MapClaims
    signature string
}

func parseIdToken(jwtString string, correctNonce string) (*idToken, error) {
    token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error){
        return []byte("base"), nil
    })
    if token.Valid {
        return nil, err
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if ok{
        if claims["nonce"] != correctNonce{
            return nil, errors.New("invalid nonce parameter")
        }
    }

    return &idToken{
        header: nil,
        claims: claims,
        signature: "",
    }, nil
}

func validateNonce(validateNonce string, correctNonce string) error {
    if validateNonce != correctNonce {
        return errors.New("invalid nonce parameter")
    }
    return nil
}


