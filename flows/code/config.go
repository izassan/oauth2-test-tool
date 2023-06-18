package code

type AuthorizeCodeFlowConfig struct{
    AuthURI string
    TokenURI string
    JwkURI string
    ClientId string
    ClientSecret string
    Scope string
    UseBrowser bool
    RequiredVerify bool
    RPConfig *RPConfig
}

type RPConfig struct {
    Host string
    Port int
}
