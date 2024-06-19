package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	logfile, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		fmt.Println("Failed to process, see logs for detail")
		return
	}

	defer logfile.Close()
	logger := slog.New(slog.NewTextHandler(logfile, nil))
	slog.SetDefault(logger)

	var inputFile, date string
	flag.StringVar(&inputFile, "f", "", "Provide the path to cookie data file")
	flag.StringVar(&date, "d", "", "Provide a date in the form yyyy-mm-dd")
	flag.Parse()

	cookies, err := findMostActiveCookieOnDay(inputFile, date)
	if err != nil {
		fmt.Println("Failed to process, see logs for detail")
		return
	}

	if len(cookies) == 0 {
		fmt.Printf("No cookies were found on %s", date)
		return
	}

	for _, cookie := range cookies {
		fmt.Println(cookie)
	}
}
