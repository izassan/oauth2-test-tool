package code

import (
	"net/url"
)

func generateAuthorizeURL(authURL string, authParams *authorizeParameters) (string, error){
    authorizeURL, err := url.Parse(authURL)
    if err != nil{
        return "", err
    }

    q := authorizeURL.Query()
    q.Add("client_id", authParams.clientId)
    q.Add("redirect_uri", authParams.redirectURI)
    q.Add("scope", authParams.scope)
    q.Add("response_type", authParams.responseType)

    if authParams.state != ""{
        q.Add("state", authParams.state)
    }
    if authParams.nonce != ""{
        q.Add("nonce", authParams.nonce)
    }
    if authParams.pkce != nil{
        q.Add("code_challenge", authParams.pkce.codeChallenge)
        q.Add("code_challenge_method", authParams.pkce.codeChallengeMethod)
    }

    authorizeURL.RawQuery = q.Encode()

    return authorizeURL.String(), nil
}
