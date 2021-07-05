package subc

import (
	"errors"
	"flag"
	"os"
)

// errors
var (
	ErrSubcNotExist = errors.New("Subcommand does not exist")
)

var subcommands = make(map[string]*flag.FlagSet)

// add new/access existing subcommand
func Sub(name string) (s *flag.FlagSet) {
	_, ok := subcommands[name]
	if !ok {
		// create the flagset if it doesn't exist
		subcommands[name] = flag.NewFlagSet(name, flag.ExitOnError)
	}
	return subcommands[name]
}

// concat all flag usages
func Usage() string {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	for f, n := range subcommands {
		fmt.Fprintf("\t%s\n", f.PrintDefaults())
	}
}

// parse the subcommand requested
// TODO: throw up errors from flag
func Parse() error {
	c, ok := subcommands[os.Args[1]]
	if !ok {
		return ErrSubcNotExist
	}
	return c.Parse(os.Args[2:])
}
