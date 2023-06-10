package code

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const letters = "abcdefghimnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateState() (string, error){
    state, err := generateRandomString(20)
    if err != nil{
        return "", err
    }
    return state, nil
}

func generateNonce() (string, error){
    nonce, err := generateRandomString(20)
    if err != nil{
        return "", err
    }
    return nonce, nil
}

func generatePKCE() (*pkce, error){
    codeVerifier, err := generateRandomString(30)
    if err != nil{
        return nil, nil
    }
    return &pkce{
        codeVerifier: codeVerifier,
        codeChallengeMethod: "S256",
        codeChallenge: generateSHA256CodeChallenge(codeVerifier),
    }, nil
}

func generateSHA256CodeChallenge(codeVerifier string) string{
    hash := sha256.Sum256([]byte(codeVerifier))
    return base64.RawURLEncoding.EncodeToString(hash[:])
}

func generateRandomString(digit uint32) (string, error){
    b := make([]byte, digit)
    if _, err := rand.Read(b); err != nil{
        return "", err
    }

    var result string
    for _, v := range b{
        result += string(letters[int(v)%len(letters)])
    }
    return result, nil
}