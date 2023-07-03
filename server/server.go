package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const CALLBACK_PATH = "/callback"
type CallbackServer struct{
    Addr string
    RedirectURI string
}

type CallbackRequest struct{
    Code string
    State string
}

func New(host string, port int) (*CallbackServer, error){
    addr := fmt.Sprintf("%s:%d", host, port)
    return &CallbackServer{
        Addr: addr,
        RedirectURI: fmt.Sprintf("http://%s%s", addr, CALLBACK_PATH),
    }, nil
}

func (cs *CallbackServer) Run(ctx context.Context, rc chan *http.Request){
    handler := func(w http.ResponseWriter, request *http.Request){
        io.WriteString(w, "<h1>Callback Success</h1>")
        rc <- request
    }
    mux := http.NewServeMux()
    mux.HandleFunc(CALLBACK_PATH, handler)
    server := &http.Server{
        Addr: cs.Addr,
        Handler: mux,
    }

    go server.ListenAndServe()
    select {
    case <-ctx.Done():
        if err := server.Shutdown(ctx); err != nil{
            fmt.Println("Failed Shutdown")
        }
    }
}

