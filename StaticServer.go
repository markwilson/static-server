package main

import (
    "net/http"
    "log"
    "path/filepath"
    "github.com/markwilson/static-server/flag"
    server_http "github.com/markwilson/static-server/http"
    "os"
)

func main() {
    f := flag.NewCommandFlagSet()
    f.Parse(os.Args[1:])

    bindAddress := f.BindAddress()

    log.Printf("Listening on http://%s/\n", bindAddress)

    path := filepath.Dir(*f.Directory)
    http.Handle("/", server_http.NewFileHandler(path))

    log.Fatal(http.ListenAndServe(bindAddress, nil))
}
