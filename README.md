# subc (WIP) [![GoDoc](https://godocs.io/github.com/e-zk/subc?status.svg)](https://godocs.io/github.com/e-zk/subc)

Wrapper for Go's `flag` that makes adding subcommands easy.

## usage

Each subcommand is a FlagSet accessed via `subc.Sub("<name>")`.  

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

	err := subc.Parse()
	if err == subc.ErrSubcNotExist {
		// show usage if an unknown subcommand is given
		print("Unknown subcommand\n")
		subc.Usage()
	} else if err != nil {
		panic(err)
	}
}
```

For code examples see the `examples/` directory.
