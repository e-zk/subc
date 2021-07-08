# Simple example

Say there is a hypothetical program that has two subcommands with the following arguments:

* `add` - write a number of random bytes to a file
  * `-n` specify the path of the file to write to
  * `-b` bytes to write
  * `-f` force overwrite file
* `remove` - remove a file specified
  * `-n` specify the path of the file to remove
  * `-f` force remove file

The subcommands are defined as follows:

```go
	subc.Sub("add").StringVar(&name, "n", "defaultname", "name of file to create")
	subc.Sub("add").IntVar(&byteLimit, "b", 128, "number of random bytes to write to file")
	subc.Sub("add").BoolVar(&force, "f", false, "force/overwrite file")

	subc.Sub("remove").StringVar(&name, "n", "defaultname", "name of file to remove")
	subc.Sub("remove").BoolVar(&force, "f", false, "force remove file")
```
