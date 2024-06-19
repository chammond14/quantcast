package main

import (
	"strings"
)

type FileLoader interface {
	Load(fileName string) ([]LogEntry, error)
}

func getFileLoader(filename string) FileLoader {
	if strings.Contains(filename, ".csv") {
		return &CsvFileLoader{}
	}

	return &CsvFileLoader{}
}
