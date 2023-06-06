package code

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func parseIdToken(jwtString string, correctNonce string) error {
    token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error){
        return []byte("base"), nil
    })
    if token.Valid {
        return err
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok{
        if claims["nonce"] != correctNonce{
            return errors.New("invalid nonce parameter")
        }
    }

    fmt.Println("nonce validation success")
    return nil
}

func validateNonce(validateNonce string, correctNonce string) error {
    if validateNonce != correctNonce {
        return errors.New("invalid nonce parameter")
    }
    return nil
}
