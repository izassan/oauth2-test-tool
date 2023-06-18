package main

import (
	"fmt"
	"os"

	"github.com/izassan/oidc-testing-tool/flows/code"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "example",
    Long: "command",
    RunE: func(cmd *cobra.Command, args []string) error {
        filePath, err := cmd.Flags().GetString("file")
        if err != nil{
            return err
        }

        ottConfig, err := ParseOTTConfigFromJsonFilePath(filePath)
        if err != nil{
            fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
        }
        ottFlags, err := getOTTFlags(cmd.Flags())
        if err != nil{
            return err
        }

        switch ottConfig.AuthFlow{
        case AUTHORIZECODEFLOW:
            authcodeConfig := &code.AuthorizeCodeFlowConfig{
                ClientId: ottConfig.ClientId,
                ClientSecret: ottConfig.ClientSecret,
                AuthURI: ottConfig.AuthURI,
                TokenURI: ottConfig.TokenURI,
                JwkURI: ottConfig.JwkURI,
                Scope: ottConfig.Scope,
                UseBrowser: !ottFlags.NoBrowser,
                RequiredVerify: !ottFlags.NoVerify,
                RPConfig: &code.RPConfig{
                    Host: ottFlags.RPHost,
                    Port: ottFlags.RPPort,
                },
            }
            if err := code.ExecuteAuthorizeCodeFlow(authcodeConfig); err != nil{
                fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
            }
        default:
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
    rootCmd.Flags().StringP(FILEFLAG, "f", "./config.json", "config file path")
    rootCmd.Flags().StringP(RPHOSTFLAG, "H", "localhost", "callback server host")
    rootCmd.Flags().IntP(RPPORTFLAG, "p", 8893, "callback server port")
    rootCmd.Flags().BoolP(NOBROWSERFLAG, "b", false, "not browser option")
    rootCmd.Flags().BoolP(NOVERIFYFLAG, "", false, "no verify id token")
}
