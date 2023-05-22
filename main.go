package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/spf13/cobra"
    "github.com/izassan/oauth2-testtool/flows"
    "github.com/izassan/oauth2-testtool/types"
)

func parseConfig(filePath string) (*types.OttConfig, error){
    f, err := os.Open(filePath)
    if err != nil{
        return nil, errors.New("file open error")
    }
    b := make([]byte, 1024)
    bcount, err := f.Read(b)
    if err != nil{
        return nil, errors.New("file open error")
    }


    var ottConfig *types.OttConfig
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
        filePath, err := cmd.Flags().GetString("file")
        if err != nil{
            return err
        }

        config, err := parseConfig(filePath)
        if err != nil{
            return err
        }
        if config.Auth_flow == "authorization_code" {
            flows.AuthorizationCodeFlow(config, flags)
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
    rootCmd.Flags().IntP("port", "p", 8893, "callback server port")
    rootCmd.Flags().BoolP("no-browser", "b", false, "not browser option")
}
