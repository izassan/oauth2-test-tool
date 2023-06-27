package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/izassan/oidc-testing-tool/flag"
)

const RESPONSETYPE_CODE= "code"

type Json map[string]interface{}
type Config interface {}

func New(jsonFilePath string, flags *flag.Flags) (Config, error) {
    fileBytes, err := readFile(jsonFilePath)
    if err != nil{
        return nil, err
    }

    var jsonObj Json
    if err := json.Unmarshal(fileBytes, &jsonObj); err != nil{
        return nil, errors.New("failed mapping from json")
    }
    provider, err := getProvider(jsonObj)
    if err != nil{
        return nil, err
    }

    cfg, err := getConfigByProvider(provider, fileBytes)
    if err != nil{
        return nil, err
    }

    return cfg, nil
}

func getProvider(j Json) (string, error){
    const KEY_PROVIDER = "provider"
    FailedGetProviderError := errors.New("failed get provider from json")

    provider, exist := j[KEY_PROVIDER]
    if !exist{
        return "", FailedGetProviderError
    }
    return provider.(string), nil
}

func getConfigByProvider(provider string, jsonBytes []byte) (Config, error){
    UnsupportProviderError := errors.New("unsupport provider")
    switch (provider){
    case GOOGLE_PROVIDER_NAME:
        var cfg GoogleOTTConfig
        if err := json.Unmarshal(jsonBytes, &cfg); err != nil{
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
