package main

import (
	"testing"
	"time"
)

func Test_isOnDayTrue(t *testing.T) {
	inputDate := "2018-12-09"
	testInputTime, _ := time.Parse(time.DateOnly, inputDate)

	testCookieTimestamp := "2018-12-09T10:30:00+00:00"
	testCookieTime, _ := time.Parse(time.RFC3339, testCookieTimestamp)

	logEntry := LogEntry{cookie: "testCookie", timestamp: testCookieTime}

	result := logEntry.isOnDay(testInputTime)
	if result != true {
		t.Errorf("Expected true but got %v", result)
	}
}

func Test_isOnDayFalse(t *testing.T) {
	inputDate := "2018-12-10"
	testInputTime, _ := time.Parse(time.DateOnly, inputDate)

	testCookieTimestamp := "2018-12-09T10:30:00+00:00"
	testCookieTime, _ := time.Parse(time.RFC3339, testCookieTimestamp)

	logEntry := LogEntry{cookie: "testCookie", timestamp: testCookieTime}

	result := logEntry.isOnDay(testInputTime)
	if result != false {
		t.Errorf("Expected false but got %v", result)
	}
}

func Test_isOnDayFalseOnlyDifferentMonth(t *testing.T) {
	inputDate := "2018-11-09"
	testInputTime, _ := time.Parse(time.DateOnly, inputDate)

	testCookieTimestamp := "2018-12-09T10:30:00+00:00"
	testCookieTime, _ := time.Parse(time.RFC3339, testCookieTimestamp)

	logEntry := LogEntry{cookie: "testCookie", timestamp: testCookieTime}

	result := logEntry.isOnDay(testInputTime)
	if result != false {
		t.Errorf("Expected false but got %v", result)
	}
}

func Test_isOnDayFalseOnlyDifferentYear(t *testing.T) {
	inputDate := "2017-12-09"
	testInputTime, _ := time.Parse(time.DateOnly, inputDate)

	testCookieTimestamp := "2018-12-09T10:30:00+00:00"
	testCookieTime, _ := time.Parse(time.RFC3339, testCookieTimestamp)

	logEntry := LogEntry{cookie: "testCookie", timestamp: testCookieTime}

	result := logEntry.isOnDay(testInputTime)
	if result != false {
		t.Errorf("Expected false but got %v", result)
	}
}
