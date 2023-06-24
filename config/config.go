package config

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

const RESPONSETYPE_CODE= "code"
const KEY_PROVIDER = "provider"

type Config interface {
    GetFullAuthorizeURI() (string, error)
    CreateSecurityParams() error
    GetTokenRequest() (*http.Request, error)
}

func New(jsonFilePath string) (Config, error) {
    fileBytes, err := readFile(jsonFilePath)
    if err != nil{
        return nil, err
    }

    provider, err := getProvider(fileBytes)
    if err != nil{
        return nil, err
    }
    cfg, err := getConfigByProvider(provider, fileBytes)
    if err != nil{
        return nil, err
    }
    return cfg, nil
}

func getProvider(fileBytes []byte) (string, error){
    FailedGetProviderError := errors.New("failed get provider from json")
    var rawJson map[string]string
    err := json.Unmarshal(fileBytes, &rawJson)
    if err != nil{
        return "", FailedGetProviderError
    }

    provider, exist := rawJson[KEY_PROVIDER]
    if !exist{
        return "", FailedGetProviderError
    }
    return provider, nil
}

func getConfigByProvider(provider string, fileBytes []byte) (Config, error){
    UnsupportProviderError := errors.New("unsupport provider")
    switch (provider){
    case GOOGLE_PROVIDER_NAME:
        var cfg GoogleOTTConfig
        if err := json.Unmarshal(fileBytes, &cfg); err != nil{
            return nil, err
        }
        return cfg, nil
    }
    return nil, UnsupportProviderError
}

func readFile(jsonFilePath string) ([]byte, error){
    f, err := os.Open(jsonFilePath)
    if err != nil{
        return nil, errors.New("file open error")
    }
    defer f.Close()
    b := make([]byte, 1024)
    c, err := f.Read(b)
    if err != nil{
        return nil, errors.New("file open error")
    }
    return b[:c], nil
}
