package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	op     string
	column int
}

func main() {
	// Verify and parse arguments
	op := flag.String("op", "sum", "Operation to be executed")
	column := flag.Int("col", 1, "CSV column on which to execute operation")

	flag.Parse()

	c := config{
		op:     *op,
		column: *column,
	}

	if err := run(flag.Args(), c, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileNames []string, c config, out io.Writer) error {
	return nil
}
