package main

import (
	"fmt"
	"time"
)

// Receives all read lines, parses them and finally displays the results
// Improvement: Use a new struct Displayer in charge of the display
type Recorder struct {
	cfg      *Config
	checkers []Checker
}

// Parses a line and sends the result to all checkers
func (r *Recorder) AddLine(line string) {
	request := Parse(line)
	for _, c := range r.checkers {
		c.AddRequest(request)
	}
}

// Flush checkers when display is done
func (r *Recorder) Flush() {
	for _, c := range r.checkers {
		c.Flush()
	}
}

// Main loop
func (r *Recorder) Record(c <-chan string) {
	recorderTicker := time.NewTicker(time.Second * time.Duration(r.cfg.Tick))
	for {
		select {
		case msg := <-c:
			// A line is received
			r.AddLine(msg)
		case <-recorderTicker.C:
			// End of a tick: we need to compute the results and display them
			fmt.Println("Collected data at", time.Now().Format("02/Jan/2006:15:04:05 -0700"))
			for _, c := range r.checkers {
				c.Compute()
				fmt.Println(c.DisplayString())
			}
			fmt.Println("")
			r.Flush()
		}
	}
}

func NewRecorder(cfg *Config) *Recorder {
	checkers := []Checker{
		&InvalidChecker{},
		&OkChecker{},
		NewAlerter(cfg),
	}
	return &Recorder{
		cfg:      cfg,
		checkers: checkers,
	}
}
