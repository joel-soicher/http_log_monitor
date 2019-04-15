package main

import (
	"fmt"
	"time"
)

// Receives all read lines, parses them and finally displays the results
type Recorder struct {
	cfg       *Config
	displayer Displayer
	checkers  []Checker
}

// Parses a line and sends the result to all checkers
func (r *Recorder) AddLine(line string) {
	request := Parse(line)
	for _, c := range r.checkers {
		c.AddRequest(request)
	}
}

// Flush checkers when display is done
func (r *Recorder) Do() {
	for _, c := range r.checkers {
		c.Compute()
		c.Display(r.displayer)
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
			header := fmt.Sprintf("Collected data at %s:\n", time.Now().Format("02/Jan/2006:15:04:05 -0700"))
			r.displayer.Display(header)
			r.Do()
			r.displayer.Display("-------------------------\n")
		}
	}
}

func NewRecorder(cfg *Config, displayer Displayer) *Recorder {
	checkers := []Checker{
		&InvalidChecker{},
		&OkChecker{},
		NewSection(cfg),
		NewSizeChecker(),
		NewAlerter(cfg, &AlertsImpl{}),
	}
	return &Recorder{
		cfg:       cfg,
		displayer: displayer,
		checkers:  checkers,
	}
}
