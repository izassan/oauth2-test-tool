package google

import (
	"net/url"
	"reflect"

	"github.com/izassan/oidc-testing-tool/config"
	"github.com/izassan/oidc-testing-tool/security"
	"github.com/izassan/oidc-testing-tool/server"
)

type GoogleOIDCFlow struct{
    AuthorizeURI string
    TokenURI string
    RedirectURI string `ap:"redirect_uri" tp:"redirect_uri"`
    ClientId string `ap:"client_id" tp:"client_id"`
    ClientSecret string `tp:"client_secret"`
    ResponseType string `ap:"response_type"`
    Scope string `ap:"scope"`
    JwkURI string
    SecurityParams *security.SecurityParams
    AccessType string `ap:"access_type"`
    Display string `ap:"display"`
    Hd string `ap:"hd"`
    LoginHint string `ap:"login_hint"`
    Prompt string `ap:"prompt"`
}

type GoogleAuthorizeResult struct{}

func New(gc config.GoogleOTTConfig, sp *security.SecurityParams, server *server.CallbackServer) (*GoogleOIDCFlow, error){
    return &GoogleOIDCFlow{
        AuthorizeURI: gc.Authorize.AuthorizeURI,
        TokenURI: gc.Token.TokenURI,
        RedirectURI: server.RedirectURI,
        ClientId: gc.ClientId,
        ClientSecret: gc.ClientSecret,
        ResponseType: "code",
        Scope: gc.Scope,
        JwkURI: gc.JwkURI,
        SecurityParams: sp,
        AccessType: gc.Authorize.AuthorizeOptions.AccessType,
        Display: gc.Authorize.AuthorizeOptions.Display,
        Hd: gc.Authorize.AuthorizeOptions.Hd,
        LoginHint: gc.Authorize.AuthorizeOptions.LoginHint,
        Prompt: gc.Authorize.AuthorizeOptions.Prompt,
    }, nil
}

func (gf GoogleOIDCFlow)GenerateAuthorizeEndpointURI() string{
    authorizeUri, err := url.Parse(gf.AuthorizeURI)
    if err != nil{
        return ""
    }
    q := authorizeUri.Query()

    fields := reflect.ValueOf(gf)
    for i:=0; i < fields.NumField(); i++{
        f := fields.Type().Field(i)
        v := fields.FieldByName(f.Name)
        ap, ok := f.Tag.Lookup("ap")
        if !ok || v.Kind() != reflect.String{
            continue
        }
        if v.String() == ""{
            continue
        }
        q.Add(ap, v.String())
    }

    authorizeUri.RawQuery = q.Encode()
    return authorizeUri.String()
}

func (gf GoogleOIDCFlow)ParseRequestToAuthorizeResult() string{
    return ""
}

func (gf GoogleOIDCFlow)ExchangeToken() string{
    return ""
}
