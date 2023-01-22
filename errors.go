package main

import "errors"

var (
	ErrNotNumber        = errors.New("data is not numeric")
	ErrInvalidColumn    = errors.New("invalid column number")
	ErrNoFiels          = errors.New("no input files")
	ErrInvalidOperation = errors.New("invalid operation")
)
