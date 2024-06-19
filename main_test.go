package main

import (
	"fmt"
	"os"
	"testing"
)

var testCsvFile, testInvalidFile *os.File

func TestMain(m *testing.M) {
	var err error
	testCsvFile, err = os.CreateTemp("./", "temp-*.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	testCsvFile.WriteString("fbcn5UAVanZf6UtG,2018-12-07T22:00:00+00:00\n")
	testCsvFile.WriteString("fbcn5UAVanZf6UtG,2018-12-07T23:30:00+00:00\n")
	testCsvFile.WriteString("4sMM2LxV07bPJzwf,2018-12-07T22:00:00+00:00\n")
	testCsvFile.WriteString("4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00\n")
	testCsvFile.WriteString("AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00\n")
	testCsvFile.WriteString("SAZuXPGUrfbcn5UA,2018-\n")
	testCsvFile.WriteString("AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00\n")

	testInvalidFile, err = os.CreateTemp("./", "temp-*.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	testInvalidFile.WriteString("4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00\n")
	testInvalidFile.WriteString("kasfghkja\n")

	exitVal := m.Run()

	os.Remove(testCsvFile.Name())
	os.Remove(testInvalidFile.Name())
	os.Exit(exitVal)
}
