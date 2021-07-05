package main

import (
	"github.com/e-zk/subc"
)

func main() {
	var (
		someb  bool
		somei  int
		someii int
	)

	subc.Sub("something").BoolVar(&someb, "q", true, "?")
	subc.Sub("something").IntVar(&somei, "n", 1, "??")
	subc.Sub("add").IntVar(&somei, "q", 15, "add: ?")
	subc.Sub("add").IntVar(&someii, "n", 2, "add: ??")

	err := subc.Parse()
	if err == subc.ErrSubcNotExist {
		subc.Usage()
	} else if err != nil {
		panic(err)
	}
}
