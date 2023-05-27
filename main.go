package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/izassan/oauth2-testtool/flows"
	"github.com/izassan/oauth2-testtool/types"
	"github.com/spf13/cobra"
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
            fmt.Println(err.Error())
            return err
        }

        config, err := parseConfig(filePath)
        if err != nil{
            fmt.Println(err.Error())
            return err
        }
        if config.Auth_flow == "authorization_code" {
            if err := flows.AuthorizationCodeFlow(config, flags); err != nil{
                fmt.Println(err.Error())
            }
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
