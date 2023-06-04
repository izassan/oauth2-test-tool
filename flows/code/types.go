package code

type authorizeParameters struct {
    clientId string
    scope string
    responseType string
    redirectURI string
    state string
    nonce string
    pkce *pkce
}

type pkce struct {
    codeChallengeMethod string
    codeChallenge string
    codeVerifier string
}
