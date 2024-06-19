package main

import "testing"

func Test_findMostActiveCookieOnDayBadDateInput(t *testing.T) {
	_, err := findMostActiveCookieOnDay(testCsvFile.Name(), "2018-12")
	if err == nil {
		t.Error("Expected error but got nil back")
	}

	if err != ErrParsingInputVar {
		t.Errorf("Incorrect error returned, expected %s but got %s", ErrParsingInputVar.Error(), err.Error())
	}
}

func Test_findMostActiveCookieOnDayBadFileInput(t *testing.T) {
	_, err := findMostActiveCookieOnDay("BadFileName.csv", "2018-12-09")
	if err == nil {
		t.Error("Expected error but got nil back")
	}

	if err != ErrOpeningFile {
		t.Errorf("Incorrect error returned, expected %s but got %s", ErrOpeningFile.Error(), err.Error())
	}
}

func Test_findMostActiveCookieOnDaySuccessMultipleResults(t *testing.T) {
	result, err := findMostActiveCookieOnDay(testCsvFile.Name(), "2018-12-07")
	if err != nil {
		t.Errorf("Expected error to be nil but got %s", err.Error())
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 cookies but got %d", len(result))
	}

	if result[0] != "fbcn5UAVanZf6UtG" {
		t.Errorf("Expected cookie fbcn5UAVanZf6UtG but got %s", result[0])
	}

	if result[1] != "4sMM2LxV07bPJzwf" {
		t.Errorf("Expected cookie 4sMM2LxV07bPJzwf but got %s", result[1])
	}
}

func Test_findMostActiveCookieOnDaySuccess(t *testing.T) {
	result, err := findMostActiveCookieOnDay(testCsvFile.Name(), "2018-12-09")
	if err != nil {
		t.Errorf("Expected error to be nil but got %s", err.Error())
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 cookie but got %d", len(result))
	}

	if result[0] != "AtY0laUfhglK3lC7" {
		t.Errorf("Expected cookie AtY0laUfhglK3lC7 but got %s", result[0])
	}
}
