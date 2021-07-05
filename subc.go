package subc

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// ErrSubcNotExist is returned when a subcommand that does not exist is requested.
var ErrSubcNotExist = errors.New("Subcommand does not exist")

var (
	subcommands            = make(map[string]*flag.FlagSet)
	outputWriter io.Writer = os.Stderr
)

// Prints a usage message documenting all defined subcommands and their flags.
var Usage = func() {
	fmt.Fprintf(outputWriter, "Usage of %s [", os.Args[0])
	for name, _ := range subcommands {
		fmt.Fprintf(outputWriter, "%s|", name)
	}
	fmt.Fprintf(outputWriter, "]\n")

	for name, f := range subcommands {
		fmt.Fprintf(outputWriter, "%s:\n", name)
		f.PrintDefaults()
	}
}

// Output returns the destination writer for usage and error messages.
func Output() io.Writer {
	return outputWriter
}

// Sets destination for all subcommand usage and messages. By default this is os.Stderr.
func SetOutput(output io.Writer) {
	outputWriter = output
}

// Sub adds a new subcommand or accesses and existing subcommand.
func Sub(name string) (s *flag.FlagSet) {
	_, ok := subcommands[name]
	if !ok {
		// create the flagset if it doesn't exist
		subcommands[name] = flag.NewFlagSet(name, flag.ExitOnError)
		subcommands[name].SetOutput(outputWriter)
	}
	return subcommands[name]
}

// Parse the subcommand input (os.Args[1]) and arguments.
func Parse() error {
	c, ok := subcommands[os.Args[1]]
	if !ok {
		return ErrSubcNotExist
	}

	return c.Parse(os.Args[2:])
}

// Parse the given subcommand input and arguments.
func ParseArgs(args []string) error {
	c, ok := subcommands[args[1]]
	if !ok {
		return ErrSubcNotExist
	}

	return c.Parse(args[2:])
}
