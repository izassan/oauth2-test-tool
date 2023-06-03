package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func parseConfig(filePath string) (*OttConfig, error){
    f, err := os.Open(filePath)
    if err != nil{
        return nil, errors.New("file open error")
    }
    b := make([]byte, 1024)
    bcount, err := f.Read(b)
    if err != nil{
        return nil, errors.New("file open error")
    }


    var ottConfig *OttConfig
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

        config, err := parseConfig(filePath)
        if err != nil{
            return err
        }
        if config.AuthFlow == AUTHORIZECODEFLOW {
        }else if config.AuthFlow == CLIENTCREDENTIALSFLOW{
        }else{
            return errors.New("Unsupported Authorization Flow. Fix 'auth_flow' parameter")
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
