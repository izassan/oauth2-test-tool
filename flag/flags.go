package flag

import "github.com/spf13/pflag"

// flag names
const (
    FILE_FLAG_NAME = "file"
    RPHOST_FLAG_NAME = "host"
    RPPORT_FLAG_NAME = "port"
    NO_BROWSER_FLAG_NAME = "no-browser"
    NO_VERIFY_FLAG_NAME = "no-verify"
)

type Flags struct {
    FilePath string
    RPHost string
    RPPort int
    NoBrowser bool
    NoVerify bool
}

func ParseFlags(flagset *pflag.FlagSet) (*Flags, error){
    fileFlag, err := flagset.GetString(FILE_FLAG_NAME)
    if err != nil{
        return nil, err
    }
    host, err := flagset.GetString(RPHOST_FLAG_NAME)
    if err != nil{
        return nil, err
    }
    port, err := flagset.GetInt(RPPORT_FLAG_NAME)
    if err != nil{
        return nil, err
    }
    noBroserFlag, err := flagset.GetBool(NO_BROWSER_FLAG_NAME)
    if err != nil{
        return nil, err
    }
    noVerifyFlag, err := flagset.GetBool(NO_VERIFY_FLAG_NAME)
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

