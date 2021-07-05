package subc_test

import (
	"github.com/e-zk/wslcheck"
)

func main() {
	var (
		someb  bool
		somei  int
		someii int
	)

	Subc("something").BoolVar(&someb, "q", true, "?")
	Subc("something").IntVar(&somei, "n", 1, "??")
	Subc("add").IntVar(&somei, "q", 15, "add: ?")
	Subc("add").IntVar(&someii, "n", 2, "add: ??")

	err := Parse()
	if err != nil {
		panic(err)
	}
}
