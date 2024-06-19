package main

import (
	"encoding/csv"
	"log/slog"
	"os"
	"time"
)

type CsvFileLoader struct{}

func (cfl *CsvFileLoader) Load(filePath string) ([]LogEntry, error) {
	slog.Info("Opening input file", "name", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		slog.Error("Failed to open file", "name", filePath)
		return nil, ErrOpeningFile
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		slog.Error("Unable to parse CSV file", "name", filePath, "error", err)
		return nil, ErrReadingFile
	}

	var entries []LogEntry
	for _, record := range records {
		c := record[0]
		ts := record[1]

		t, err := time.Parse(time.RFC3339, ts) // RFC3339 layout format "2006-01-02T15:04:05Z07:00"
		if err != nil {
			slog.Info("Invalid timestamp found for cookie, ignoring entry", "cookie", c)
			continue
		}

		entry := LogEntry{
			c,
			t,
		}

		slog.Info("Parsed CSV record with values", "cookie", c, "timestamp", ts)
		entries = append(entries, entry)
	}

	return entries, nil
}
