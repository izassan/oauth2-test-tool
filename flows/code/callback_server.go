package code

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func startCallbackServer(ctx context.Context, host string, port int, rc chan *http.Request){
    handler := func(w http.ResponseWriter, request *http.Request){
        io.WriteString(w, "<h1>Callback Success</h1>")
        rc <- request
    }
    mux := http.NewServeMux()
    mux.HandleFunc("/callback", handler)

    addr := fmt.Sprintf("%s:%d", host, port)
    server := &http.Server{
        Addr: addr,
        Handler: mux,
    }

    go server.ListenAndServe()
    select {
    case <-ctx.Done():
        if err := server.Shutdown(ctx); err != nil{
            fmt.Println("Failed Shutdown")
            os.Exit(1)
        }
    }
}

