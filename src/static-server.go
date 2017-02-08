package main

import (
    "net/http"
    "log"
    "path/filepath"
    "flag"
    "strconv"
    "fmt"
)

func main() {
    portPointer   := flag.Int("p", 8080, "The port to listen on")
    exposePointer := flag.Bool("x", false, "Expose the server outside of localhost")

    flag.Parse()

    var bindAddress string

    if *exposePointer {
        bindAddress = "0.0.0.0:" + strconv.Itoa(*portPointer)
    } else {
        bindAddress = "127.0.0.1:" + strconv.Itoa(*portPointer)
    }

    fmt.Printf("Listening on http://%s/\n", bindAddress)

    path := filepath.Dir(".")
    http.Handle("/", http.FileServer(http.Dir(path)))

    log.Fatal(http.ListenAndServe(bindAddress, nil))
}
