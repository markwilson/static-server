package flag

import (
	"os"
	"flag"
	"errors"
	"strconv"
)

type CommandFlagSet struct {
	*flag.FlagSet
	Port      *int
	Expose    *bool
	Directory *string
	Ip        *IpAddressFlag
}

func (f *CommandFlagSet) Validate() error {
	if f.Ip.set && *f.Expose {
		// TODO: add usage & invalid message to output
		f.PrintDefaults()
		return errors.New("Invalid flags")
	}

	if *f.Expose {
		f.Ip.value = "0.0.0.0"
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

func (f *CommandFlagSet) BindAddress() string {
	return (*f.Ip).String() + ":" + strconv.Itoa(*f.Port)
}

func NewCommandFlagSet() CommandFlagSet {
	f := CommandFlagSet{
		FlagSet: flag.NewFlagSet(os.Args[0], flag.ExitOnError),
	}

	f.Port = f.Int("p", 8080, "The Port to listen on")
	f.Expose = f.Bool("x", false, "Shorthand for -i 0.0.0.0")
	f.Directory = f.String("d", ".", "Path of static files to host")

	var bindAddress IpAddressFlag
	f.Var(&bindAddress, "i", "Specify a bind address - can not be used in conjunction with the -x flag")
	bindAddress.value = "127.0.0.1"
	f.Ip = &bindAddress

	return f
}
