package main

import (
    "net/http"
    "log"
    "path/filepath"
    "flag"
    "strconv"
    "os"
    "errors"
)

type FileHandler struct {
    http.Handler
}

func (h FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Printf("Request for %s", r.RequestURI)
    h.Handler.ServeHTTP(w, r)
}

type IpAddressFlag struct {
    set bool
    value string
}

func (f *IpAddressFlag) Set(ip string) error {
    f.value = ip
    f.set = true

    return nil
}

func (f *IpAddressFlag) String() string {
    return f.value
}

type CommandFlagSet struct {
    *flag.FlagSet
    port *int
    expose *bool
    directory *string
    bindAddress *IpAddressFlag
}

func (f *CommandFlagSet) Validate() error {
    if f.bindAddress.set && *f.expose {
        // TODO: add usage & invalid message to output
        f.PrintDefaults()
        return errors.New("Invalid flags")
    }

    if *f.expose {
        f.bindAddress.value = "0.0.0.0"
    }

    return nil
}

func (f *CommandFlagSet) Parse(arguments []string) error {
    f.FlagSet.Parse(arguments)

    err := f.Validate()
    if err != nil {
        os.Exit(2)
    }

    return nil
}

func NewCommandFlagSet() CommandFlagSet {
    f := CommandFlagSet{
        FlagSet: flag.NewFlagSet(os.Args[0], flag.ExitOnError),
    }

    f.port      = f.Int("p", 8080, "The port to listen on")
    f.expose    = f.Bool("x", false, "Shorthand for -i 0.0.0.0")
    f.directory = f.String("d", ".", "Path of static files to host")

    var bindAddress IpAddressFlag
    f.Var(&bindAddress, "i", "Specify a bind address - can not be used in conjunction with the -x flag")
    bindAddress.value = "127.0.0.1"
    f.bindAddress = &bindAddress

    return f
}

func main() {
    f := NewCommandFlagSet()
    f.Parse(os.Args[1:])

    bindAddress := (*f.bindAddress).String() + ":" + strconv.Itoa(*f.port)

    log.Printf("Listening on http://%s/\n", bindAddress)

    path := filepath.Dir(*f.directory)
    http.Handle("/", FileHandler{Handler: http.FileServer(http.Dir(path))})

    log.Fatal(http.ListenAndServe(bindAddress, nil))
}
