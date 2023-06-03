package code

func generateState() string{
    return "dummy-state"
}

func generateNonce() string{
    return "dummy-nonce"
}

func generatePKCE() pkce{
    return pkce{
        codeChallenge: "dummy-codechallenge",
        codeChallengeMethod: "SHA256",
    }
}
