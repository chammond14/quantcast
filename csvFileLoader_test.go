package main

import (
	"testing"
)

func TestLoadMissingFile(t *testing.T) {
	csvLoader := CsvFileLoader{}
	_, err := csvLoader.Load("missingFile.csv")
	if err == nil {
		t.Errorf("Expected error but got nil back")
	}
	if err != ErrOpeningFile {
		t.Errorf("Incorrect error returned, expected %s but got %s", ErrOpeningFile.Error(), err.Error())
	}
}

func TestLoadInvalidCsv(t *testing.T) {
	csvLoader := CsvFileLoader{}
	_, err := csvLoader.Load(testInvalidFile.Name())
	if err == nil {
		t.Errorf("Expected error but got nil back")
	}
	if err != ErrReadingFile {
		t.Errorf("Incorrect error returned, expected %s but got %s", ErrReadingFile.Error(), err.Error())
	}
}

func TestLoadInvalidTimestampOmitsCookie(t *testing.T) {
	csvLoader := CsvFileLoader{}
	entries, err := csvLoader.Load(testCsvFile.Name())
	if err != nil {
		t.Errorf("Expected error to be nil but got %s", err.Error())
	}

	cookieWithMissingEntry := "SAZuXPGUrfbcn5UA"
	var cookieCount int

	for _, entry := range entries {
		if entry.cookie == cookieWithMissingEntry {
			cookieCount++
		}
	}

	if cookieCount != 0 {
		t.Errorf("Expected cookieCount to be 0 but got %d", cookieCount)
	}
}

func TestLoadReturnsCookieEntries(t *testing.T) {
	csvLoader := CsvFileLoader{}
	entries, err := csvLoader.Load(testCsvFile.Name())
	if err != nil {
		t.Errorf("Expected error to be nil but got %s", err.Error())
	}

	if len(entries) != 6 {
		t.Errorf("Expected 6 entries to be loaded, got %d", len(entries))
	}
}
