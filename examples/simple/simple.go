package main

import (
	"go.zakaria.org/subc"
)

func main() {
	var (
		force     bool
		name      string
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
	} else if err == subc.ErrUsage {
		return
	} else if err != nil {
		panic(err)
	}

	print(subcommand + "\n")
}
