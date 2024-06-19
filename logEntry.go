package main

import (
	"time"
)

type LogEntry struct {
	cookie    string
	timestamp time.Time
}

func (le LogEntry) isOnDay(t time.Time) bool {
	return t.Day() == le.timestamp.Day() && t.Month() == le.timestamp.Month() && t.Year() == le.timestamp.Year()
}
