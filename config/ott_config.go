package config

type OttConfig struct{
    AuthFlow string `json:"auth_flow"`
    AuthURI string `json:"auth_uri"`
    TokenURI string `json:"token_uri"`
    ClientId string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
    Scope string `json:"scope"`
}
