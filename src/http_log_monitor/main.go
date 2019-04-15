package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Parse command line arguments
	file := flag.String("file", "/tmp/access.log", "path to access log file")
	tick := flag.Int("tick", 10, "number of seconds between two displays")
	alertDelay := flag.Int("alert", 12, "number of ticks in which alert can be triggered")
	maxReq := flag.Int("maxreq", 10, "max number of request per second")
	maxSections := flag.Int("maxsections", 10, "Maximum number of most hit sections to be displayed")

	flag.Parse()

	if file == nil || !fileExists(*file) {
		panic("input file is not valid")
	}

	if tick == nil {
		panic("invalid tick param")
	}

	if alertDelay == nil {
		panic("invalid alert param")
	}

	if maxReq == nil {
		panic("invalid maxreq param")
	}

	if maxSections == nil {
		panic("invalid maxsections param")
	}

	cfg := &Config{
		Tick:        *tick,
		AlertDelay:  *alertDelay,
		MaxReq:      *maxReq,
		MaxSections: *maxSections,
	}

	// Channel to exit cleanly
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("captured %v, stopping and exiting\n", sig)
			os.Exit(1)
		}
	}()

	// Channel receiving the read lines
	msgChan := make(chan string)

	// Create the console Displayer
	displayer := &ConsoleDisplayer{}

	// Creates the recorder and launch it in a go routine
	recorder := NewRecorder(cfg, displayer)
	go recorder.Record(msgChan)

	// Reads the file
	consume(*file, msgChan)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
