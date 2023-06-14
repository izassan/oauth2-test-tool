package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/izassan/oidc-testing-tool/config"
	"github.com/izassan/oidc-testing-tool/flows/code"
	"github.com/spf13/cobra"
)

func parseConfig(filePath string) (*config.OttConfig, error){
    f, err := os.Open(filePath)
    if err != nil{
        return nil, errors.New("file open error")
    }
    b := make([]byte, 1024)
    bcount, err := f.Read(b)
    if err != nil{
        return nil, errors.New("file open error")
    }


    var ottConfig *config.OttConfig
    if err := json.Unmarshal(b[:bcount], &ottConfig); err != nil{
        return nil, errors.New("file open error")

    }
    return ottConfig, nil
}

var rootCmd = &cobra.Command{
    Use: "example",
    Long: "command",
    RunE: func(cmd *cobra.Command, args []string) error {
        flags := cmd.Flags()
        filePath, err := flags.GetString("file")
        if err != nil{
            return err
        }

        ottConfig, err := parseConfig(filePath)
        if err != nil{
            fmt.Printf("error: %s\n", err.Error())
            return nil
        }
        if ottConfig.AuthFlow == AUTHORIZECODEFLOW {
            if err := code.ExecuteAuthorizeCodeFlow(ottConfig, flags); err != nil{
                fmt.Printf("error: %s\n", err.Error())
                return nil
            }
        }else if ottConfig.AuthFlow == CLIENTCREDENTIALSFLOW{
        }else{
            fmt.Println("Unsupported Authorization Flow. Fix 'auth_flow' parameter")
        }
        return nil
    },
}

func main(){
    err := rootCmd.Execute()
    if err != nil{
        os.Exit(1)
    }
}

func init() {
    rootCmd.Flags().StringP("file", "f", "./config.json", "config file path")
    rootCmd.Flags().StringP("host", "H", "localhost", "callback server host")
    rootCmd.Flags().IntP("port", "p", 8893, "callback server port")
    rootCmd.Flags().BoolP("no-browser", "b", false, "not browser option")
    rootCmd.Flags().BoolP("no-verify", "", false, "no verify id token")
}
