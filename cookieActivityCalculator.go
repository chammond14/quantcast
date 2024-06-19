package main

import (
	"fmt"
	"log/slog"
	"time"
)

func findMostActiveCookieOnDay(filePath string, date string) ([]string, error) {
	slog.Info("Searching for most active cookie:", "file", filePath, "date", date)

	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		slog.Error("Error parsing input date", "date", date, "err", err)
		fmt.Println("Invalid date provided, expected format yyyy-mm-dd")
		return nil, ErrParsingInputVar
	}

	loader := getFileLoader(filePath)
	logEntries, err := loader.Load(filePath)
	if err != nil {
		return nil, err
	}

	cookieOccurrenceMap := make(map[string]int)
	for _, entry := range logEntries {
		if entry.isOnDay(d) {
			slog.Info("Cookie activity matches date filter", "cookie", entry.cookie)
			cookieOccurrenceMap[entry.cookie]++
		}
	}

	var mostActiveCookies []string
	var mostActiveCount int

	for cookie, count := range cookieOccurrenceMap {
		if count > mostActiveCount {
			slog.Info("Setting new cookie as current most active", "cookie", cookie, "count", count)
			mostActiveCookies = nil
			mostActiveCookies = append(mostActiveCookies, cookie)
			mostActiveCount = count
		} else if count == mostActiveCount {
			slog.Info("Cookie has equal count to the current most active", "cookie", cookie, "count", count)
			mostActiveCookies = append(mostActiveCookies, cookie)
		}
	}

	slog.Info("Processing complete", "most active cookie(s)", mostActiveCookies)
	return mostActiveCookies, nil
}
