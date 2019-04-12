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

	cfg := &Config{
		Tick:       *tick,
		AlertDelay: *alertDelay,
		MaxReq:     *maxReq,
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

	// Creates the recorder and launch it in a go routine
	recorder := NewRecorder(cfg)
	go recorder.Record(msgChan)

	// Reads the file
	consume(*file, msgChan)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
