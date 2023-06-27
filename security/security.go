package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const letters = "abcdefghimnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type SecurityParams struct{
    Nonce string
    State string
    Pkce *pkce
}

type pkce struct {
    CodeChallengeMethod string
    CodeChallenge string
    CodeVerifier string
}

type (
    CodeSecurityConfig struct{
        WithoutNonce bool
        WithoutState bool
        WithoutPKCE bool
    }
    codeSecurityOption func(*CodeSecurityConfig)error
)

func withoutNonce() codeSecurityOption{
    return func(c *CodeSecurityConfig)error{
        c.WithoutNonce = true
        return nil
    }
}
func withoutState() codeSecurityOption{
    return func(c *CodeSecurityConfig)error{
        c.WithoutState = true
        return nil
    }
}
func withoutPKCE() codeSecurityOption{
    return func(c *CodeSecurityConfig)error{
        c.WithoutPKCE = true
        return nil
    }
}


func NewSecurityParams(options ...codeSecurityOption) (*SecurityParams, error){
    config := &CodeSecurityConfig{
        WithoutNonce: false,
        WithoutState: false,
        WithoutPKCE: false,
    }

    for _, opt := range options{
        opt(config)
    }

    nonce, err := generateRandomString(20, config.WithoutNonce)
    if err != nil{
        return nil, err
    }
    state, err := generateRandomString(20, config.WithoutState)
    if err != nil{
        return nil, err
    }
    pkce, err := generatePKCE(config.WithoutPKCE)
    if err != nil{
        return nil, err
    }

    return &SecurityParams{
        Nonce: nonce,
        State: state,
        Pkce: pkce,
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
        CodeVerifier: codeVerifier,
        CodeChallengeMethod: "S256",
        CodeChallenge: generateSHA256CodeChallenge(codeVerifier),
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
