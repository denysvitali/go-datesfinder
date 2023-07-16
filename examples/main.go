package main

import (
	"fmt"
	"github.com/denysvitali/go-datesfinder"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	text, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	dates, _ := datesfinder.FindDates(string(text))
	fmt.Printf("Found %d dates:\n", len(dates))
	for _, date := range dates {
		fmt.Printf("- %s\n", date.Format("2006-01-02"))
	}
}
