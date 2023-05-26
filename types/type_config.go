package types

type OttConfig struct{
    Provider string `json:"provider"`
    Auth_flow string `json:"auth_flow"`
    Auth_uri string `json:"auth_uri"`
    Token_uri string `json:"token_uri"`
    Client_id string `json:"client_id"`
    Client_secret string `json:"client_secret"`
    CredentialFile string `json:"credential_file"`
    Scope string `json:"scope"`
    Require_string string `json:"require_string"`
}
