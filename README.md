# subc [![GoDoc](https://godocs.io/github.com/e-zk/subc?status.svg)](https://godocs.io/github.com/e-zk/subc)

Wrapper for Go's `flag` that makes adding subcommands easy.

## Usage

Each subcommand is a [`flag.FlagSet`](https://godocs.io/flag) accessed via `subc.Sub("<name>")`, so you can do anything you would be able to do with a normal Go `flag.FlagSet`.

### Example

For example say there is a hypothetical program that has two subcommands:

* `add` - write a number of random bytes (specified by `-b`) to a file (specified by `-n`) with an option to force overwrite (`-f`).
* `remove` - remove the file specified by (`-n`) with an option to force remove (`-f`).

The subcommands can be defined as follows:

```go
package main

import (
        "github.com/e-zk/subc"
)

func main() {
	var (
		force bool
		name string
		byteLimit int
	)

	subc.Sub("add").StringVar(&name, "n", "defaultname", "name of file to create")
	subc.Sub("add").IntVar(&byteLimit, "b", 128, "number of random bytes to write to file")
	subc.Sub("add").BoolVar(&force, "f", false, "force/overwrite file")

	subc.Sub("remove").StringVar(&name, "n", "defaultname", "name of file to remove")
	subc.Sub("remove").BoolVar(&force, "f", false, "force remove file")

	subcommand, err := subc.Parse()
	if err == subc.ErrSubcNotExist {
		// show usage if an unknown subcommand is given
		print("Unknown subcommand\n")
		subc.Usage()
	} else if err != nil {
		panic(err)
	}

	print(subcommand + "\n")
}
```

For code examples see the `examples/` directory.
