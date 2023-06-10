package code

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const letters = "abcdefghimnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type securityParams struct{
    nonce string
    state string
    pkce *pkce
}

type pkce struct {
    codeChallengeMethod string
    codeChallenge string
    codeVerifier string
}

type (
    codeSecurityConfig struct{
        withoutNonce bool
        withoutState bool
        withoutPKCE bool
    }
    codeSecurityOption func(*codeSecurityConfig)error
)

func withoutNonce() codeSecurityOption{
    return func(c *codeSecurityConfig)error{
        c.withoutNonce = true
        return nil
    }
}
func withoutState() codeSecurityOption{
    return func(c *codeSecurityConfig)error{
        c.withoutState = true
        return nil
    }
}
func withoutPKCE() codeSecurityOption{
    return func(c *codeSecurityConfig)error{
        c.withoutPKCE = true
        return nil
    }
}


func newSecurityParams(options ...codeSecurityOption) (*securityParams, error){
    config := &codeSecurityConfig{
        withoutNonce: false,
        withoutState: false,
        withoutPKCE: false,
    }

    for _, opt := range options{
        opt(config)
    }

    nonce, err := generateRandomString(20, config.withoutNonce)
    if err != nil{
        return nil, err
    }
    state, err := generateRandomString(20, config.withoutState)
    if err != nil{
        return nil, err
    }
    pkce, err := generatePKCE(config.withoutPKCE)
    if err != nil{
        return nil, err
    }

    return &securityParams{
        nonce: nonce,
        state: state,
        pkce: pkce,
    }, nil
}

func generatePKCE(disabled bool) (*pkce, error){
    if disabled{
        return nil, nil
    }
    codeVerifier, err := generateRandomString(30, false)
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

func generateRandomString(digit uint32, disabled bool) (string, error){
    if disabled{
        return "", nil
    }
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
