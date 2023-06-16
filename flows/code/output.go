package code

import "fmt"

func outputResult(t *token, it *idToken, formatType string){
    if formatType != "default"{
        fmt.Println("warning: specified unsupported or unimplement output format. output default format")
    }
    outputTokenInfo(t)
    if it != nil{
        fmt.Printf("----------------------\n")
        outputIdTokenInfo(it)
    }
}

func outputTokenInfo(t *token) {
	fmt.Printf("access_token: %s\n", t.AccessToken)
	fmt.Printf("id_token: %s\n", t.IdToken)
	fmt.Printf("refresh_token: %s\n", t.RefreshToken)
	fmt.Printf("scope: %s\n", t.Scope)
	fmt.Printf("token_type: %s\n", t.TokenType)
	fmt.Printf("expire_in: %d\n", t.ExpiresIn)
}

func outputIdTokenInfo(t *idToken){
        fmt.Printf("id_token audiences: %s\n", t.Audience())
        fmt.Printf("id_token issuer: %s\n", t.Issuer())
}
