package subc

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// ErrSubcNotExist is returned when a subcommand that does not exist is
// requested.
var ErrSubcNotExist = errors.New("Subcommand does not exist")

// ErrNoSubc is returned when the given arguments do not contain a subcommand.
var ErrNoSubc = errors.New("Subcommand not given")

// ErrUsage is returned when usage/help is requested.
var ErrUsage = errors.New("Usage requested")

var (
	subcommands            = make(map[string]*flag.FlagSet)
	outputWriter io.Writer = os.Stderr
)

// isHelp determines whether the given string is requesting the help/usage
// message.
func isHelp(subcommand string) bool {
	switch subcommand {
	case "help":
		fallthrough
	case "-help":
		fallthrough
	case "h":
		fallthrough
	case "-h":
		return true
	}

	return false
}

// Prints a usage message documenting all defined subcommands and their flags.
var Usage = func() {
	// create an array of subcommand names
	names := make([]string, len(subcommands))
	i := 0
	for name := range subcommands {
		names[i] = name
		i++
	}

	// print binary name, subcommand list
	fmt.Fprintf(outputWriter, "Usage of %s [", os.Args[0])
	fmt.Fprintf(outputWriter, "%s", strings.Join(names, "|"))
	fmt.Fprintf(outputWriter, "]\n")

	// for each subcommand print it's usage
	for name, f := range subcommands {
		fmt.Fprintf(outputWriter, "%s:\n", name)
		f.PrintDefaults()
	}
}

// Output returns the destination writer for usage and error messages.
func Output() io.Writer {
	return outputWriter
}

// Sets destination for all subcommand usage and messages. By default this is
// os.Stderr.
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

// Parse parses the flags for the subcommand os.Args[1]. Must be called after
// all flags are defined and before flags are accessed by the program. Returns
// the name of the requested subcommand.
func Parse() (string, error) {
	if !(len(os.Args) > 1) {
		return "", ErrNoSubc
	}

	if isHelp(os.Args[1]) {
		Usage()
		return "", ErrUsage
	}

	c, ok := subcommands[os.Args[1]]
	if !ok {
		return "", ErrSubcNotExist
	}

	return os.Args[1], c.Parse(os.Args[2:])
}

// Parse parses the the flags of the subcommand given as the first item in the
// argument list. The given argument list should not include the command name.
// Returns the name of the parsed subcommand.
func ParseArgs(arguments []string) (string, error) {
	if !(len(arguments) > 0) {
		return "", ErrNoSubc
	}

	if isHelp(arguments[0]) {
		Usage()
		return "", ErrUsage
	}
	c, ok := subcommands[arguments[0]]
	if !ok {
		return "", ErrSubcNotExist
	}

	return arguments[1], c.Parse(arguments[2:])
}
