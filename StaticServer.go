package main

import (
    "net/http"
    "log"
    "github.com/markwilson/static-server/flag"
    server_http "github.com/markwilson/static-server/http"
    "os"
)

func main() {
    f := flag.NewCommandFlagSet()
    f.Parse(os.Args[1:])

    bindAddress := f.BindAddress()
    path        := *f.Directory

    log.Printf("Listening on http://%s/ in %s\n", bindAddress, path)

    http.Handle("/", server_http.NewFileHandler(path))
    log.Fatal(http.ListenAndServe(bindAddress, nil))
}
