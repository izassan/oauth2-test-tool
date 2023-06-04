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

type authorizeCodeParameters struct {
    code string
    state string
    scope string
}

type tokenExchangeParams struct {
    code string
    clientId string
    clientSecret string
    redirectURI string
    grantType string
    codeVerifier string
}

type token struct {
    accessToken string
    expiresIn int
    idToken string
    scope string
    tokenType string
    refreshToken string
}
