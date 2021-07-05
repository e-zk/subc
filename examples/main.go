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

	subc.Subc("something").BoolVar(&someb, "q", true, "?")
	subc.Subc("something").IntVar(&somei, "n", 1, "??")
	subc.Subc("add").IntVar(&somei, "q", 15, "add: ?")
	subc.Subc("add").IntVar(&someii, "n", 2, "add: ??")

	err := subc.Parse()
	if err != nil {
		panic(err)
	}
}
