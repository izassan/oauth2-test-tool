package code

import "fmt"

func outputTokenInfo(t *token) {
	fmt.Printf("access_token: %s\n", t.AccessToken)
	fmt.Printf("id_token: %s\n", t.IdToken)
	fmt.Printf("refresh_token: %s\n", t.RefreshToken)
	fmt.Printf("scope: %s\n", t.Scope)
	fmt.Printf("token_type: %s\n", t.TokenType)
	fmt.Printf("expire_in: %d\n", t.ExpiresIn)
}
