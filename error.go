package main

import "errors"

var (
	ErrOpeningFile     = errors.New("could not open file")
	ErrReadingFile     = errors.New("could not read input file")
	ErrParsingInputVar = errors.New("could not parse input variable")
)
