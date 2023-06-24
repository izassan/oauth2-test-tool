package main

import (
	"fmt"
	"os"

	"github.com/izassan/oidc-testing-tool/config"
	"github.com/izassan/oidc-testing-tool/flag"
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

        cfg, err := config.New(filePath)
        if err != nil{
            fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
        }

        if err := cfg.CreateSecurityParams(); err != nil{
            return err
        }

        // TODO: pass to flow
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
    rootCmd.Flags().StringP(flag.FILE_FLAG_NAME, "f", "./config.json", "config file path")
    rootCmd.Flags().StringP(flag.RPHOST_FLAG_NAME, "H", "localhost", "callback server host")
    rootCmd.Flags().IntP(flag.RPPORT_FLAG_NAME, "p", 8893, "callback server port")
    rootCmd.Flags().BoolP(flag.NO_BROWSER_FLAG_NAME, "b", false, "not browser option")
    rootCmd.Flags().BoolP(flag.NO_VERIFY_FLAG_NAME, "", false, "no verify id token")
}
