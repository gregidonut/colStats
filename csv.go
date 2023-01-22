package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

func sum(data []float64) float64 {
	var sum float64
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// statsFunc is an auxiliary type using the same signature
// as the above sum() and avg() this will prove useful when
// this type is used as an input parameter on a calling
// to better manage testing and make the calling function more
// concise
type statsFunc func(data []float64) float64

// csv2Float parses contents of the csv file into a slice of floats
// to perform calculations on
func csv2Float(r io.Reader, column int) ([]float64, error) {
	// Create the CSV Reader used to read in data from CSV files
	cr := csv.NewReader(r)

	// Adjusting for 0 based index, the program assumes the users
	// will input columns starting from 1, as it's more natural to understand
	column--

	// read in all CSV data
	allData, err := cr.ReadAll()
	if err != nil {
		// Wrapping error with %w verb to decorate error with
		// additional information
		return nil, fmt.Errorf("cannot read data from file: %w", err)
	}

	// convert data from csv into a slice of floats to perform calculations on
	data := make([]float64, 0)
	// looping through all records
	for i, row := range allData {
		// ignoring first record since this would be the column headers
		if i == 0 {
			continue
		}
		// checking number of columns in CSV file
		// if file doesn't have the column number specified
		if len(row) <= column {
			fmt.Errorf("%w: file has only %d columns", ErrInvalidColumn, len(row))
		}
		// try to convert data into a float
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}

	return data, nil
}