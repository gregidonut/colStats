package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
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

	// Create the channel to receive results or errors of operations
	resultsCh := make(chan []float64)
	errorCh := make(chan error)
	doneCh := make(chan struct{})

	wg := sync.WaitGroup{}

	// loop through data concurrently to be consolidated as soon as
	// resultsCh is written too which should mean that there were
	// no errors in parsing the file and writing the data to the resultsCh
	for _, fName := range fileNames {
		wg.Add(1)

		go func(fName string) {
			defer wg.Done()

			// open file for reading
			f, err := os.Open(fName)
			if err != nil {
				errorCh <- fmt.Errorf("cannot open file: %w", err)
				return
			}

			// parse the csv into a slice of float64
			data, err := csv2Float(f, cfg.column)
			if err != nil {
				errorCh <- err
			}

			if err := f.Close(); err != nil {
				errorCh <- err
			}

			resultsCh <- data
		}(fName)
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errorCh:
			return err
		case data := <-resultsCh:
			consolidate = append(consolidate, data...)
		case <-doneCh:
			_, err := fmt.Fprintln(out, opFunc(consolidate))
			return err
		}
	}
}
