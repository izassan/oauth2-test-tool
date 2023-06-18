package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

// Authorize Flow Names
const (
    AUTHORIZECODEFLOW = "code"
)

type OTTConfig struct{
    AuthFlow string `json:"auth_flow"`
    AuthURI string `json:"auth_uri"`
    TokenURI string `json:"token_uri"`
    ClientId string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
    Scope string `json:"scope"`
    JwkURI string `json:"jwk_uri"`
}

func ParseOTTConfigFromJsonFilePath(fp string) (*OTTConfig, error){
        jsonBytes, err := ReadJson(fp)
        if err != nil{
            return nil, err
        }
        ottConfig, err := ParseOTTConfigFromJson(jsonBytes)
        if err != nil{
            return nil, err
        }
        return ottConfig, err
}

func ParseOTTConfigFromJson(jsonBytes []byte) (*OTTConfig, error){
    var ottConfig OTTConfig
    if err := json.Unmarshal(jsonBytes, &ottConfig); err != nil{
        log.Println(string(jsonBytes))
        return nil, errors.New("failed to parse file to json")
    }
    return &ottConfig, nil
}

func ReadJson(jsonFilePath string) ([]byte, error){
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
