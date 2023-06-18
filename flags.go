package main

import "github.com/spf13/pflag"

// flag names
const (
    FILEFLAG = "file"
    RPHOSTFLAG = "host"
    RPPORTFLAG = "port"
    NOBROWSERFLAG = "no-browser"
    NOVERIFYFLAG = "no-verify"
)

type Flags struct {
    FilePath string
    RPHost string
    RPPort int
    NoBrowser bool
    NoVerify bool
}

func getOTTFlags(flagset *pflag.FlagSet) (*Flags, error){
    fileFlag, err := flagset.GetString(FILEFLAG)
    if err != nil{
        return nil, err
    }
    host, err := flagset.GetString(RPHOSTFLAG)
    if err != nil{
        return nil, err
    }
    port, err := flagset.GetInt(RPPORTFLAG)
    if err != nil{
        return nil, err
    }
    noBroserFlag, err := flagset.GetBool(NOBROWSERFLAG)
    if err != nil{
        return nil, err
    }
    noVerifyFlag, err := flagset.GetBool(NOVERIFYFLAG)
    if err != nil{
        return nil, err
    }

    return &Flags{
        FilePath: fileFlag,
        RPHost: host,
        RPPort: port,
        NoBrowser: noBroserFlag,
        NoVerify: noVerifyFlag,
    }, nil
}

