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

func run(fileNames []string, cfg config, out io.Writer) error {
	var opFunc statsFunc

	// Validate user inputs as a QOL for the user
	if len(fileNames) == 0 {
		return ErrNoFiles
	}
	if cfg.column < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, cfg.column)
	}
	switch cfg.op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, cfg.op)
	}

	// process CSV files
	consolidate := make([]float64, 0)
	// loop through all files adding their data to consolidate
	for _, fName := range fileNames {
		// open file for reading
		f, err := os.Open(fName)
		if err != nil {
			return fmt.Errorf("cannot open file: %w", err)
		}

		// parse the csv into a slice of float64
		data, err := csv2Float(f, cfg.column)
		if err != nil {
			return err
		}

		// Append the data to consolidate
		consolidate = append(consolidate, data...)
	}
	_, err := fmt.Fprintln(out, opFunc(consolidate))
	return err
}