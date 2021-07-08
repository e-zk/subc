# subc [![GoDoc](https://godocs.io/github.com/e-zk/subc?status.svg)](https://godocs.io/github.com/e-zk/subc)

Wrapper for Go's `flag` that makes adding subcommands easy.

## Usage

Each subcommand is a [`flag.FlagSet`](https://godocs.io/flag) accessed via `subc.Sub("<name>")`, so you can do anything you would be able to do with a normal Go `flag.FlagSet`.

For code samples see the `examples/` directory.
