package main

import (
	"fmt"
	"time"
)

// Defines a persistent alert event
type Alert interface {
	DisplayAlert() string
}

type Alerts interface {
	Trigger(hits float64)
	Untrigger()
	DisplayAlerts() string
}

// Defines a alert activation event (extends Alert)
type AlertActivation struct {
	hits float64
	date string
}

func (a *AlertActivation) DisplayAlert() string {
	return fmt.Sprintf("High traffic generated an alert - hits = %.2f requests/s, triggered at %s\n", a.hits, a.date)
}

// Defines a alert deactivation event (extends Alert)
type AlertDeactivation struct {
	date string
}

func (a *AlertDeactivation) DisplayAlert() string {
	return fmt.Sprintf("Alert recovered at %s\n", a.date)
}

// Contains the alerts and implement Alerts interface
type AlertsImpl struct {
	alerts []Alert
}

func (a *AlertsImpl) Trigger(hits float64) {
	a.alerts = append(a.alerts, &AlertActivation{hits: hits, date: time.Now().Format("02/Jan/2006:15:04:05 -0700")})
}

func (a *AlertsImpl) Untrigger() {
	a.alerts = append(a.alerts, &AlertDeactivation{date: time.Now().Format("02/Jan/2006:15:04:05 -0700")})
}

func (a *AlertsImpl) DisplayAlerts() string {
	var ret string
	for _, alert := range a.alerts {
		ret += alert.DisplayAlert()
	}
	return ret
}

// Checks if alerts must be triggered or untriggered (extends Checker)
type Alerter struct {
	cfg       *Config
	alerts    Alerts
	counts    []int
	index     int
	alerted   bool
	firstPass bool
}

func NewAlerter(cfg *Config, alerts Alerts) *Alerter {
	return &Alerter{
		cfg:       cfg,
		alerts:    alerts,
		counts:    make([]int, cfg.AlertDelay),
		firstPass: true,
	}
}

// Increments the number of request per tick
func (a *Alerter) AddRequest(req *Request) {
	a.counts[a.index]++
}

// Detects if an alert is triggered or untriggered.
func (a *Alerter) Compute() {
	avg := a.computeAverage()
	if avg >= float64(a.cfg.MaxReq) {
		if !a.alerted {
			a.alerts.Trigger(avg)
		}
		a.alerted = true
	} else {
		if a.alerted {
			a.alerts.Untrigger()

		}
		a.alerted = false
	}
	a.updateIndex()
}

// Updates and resets the rotating array containing the number of requests per tick
func (a *Alerter) updateIndex() {
	a.index++
	if a.index >= len(a.counts) {
		a.firstPass = false
		a.index = 0
	}
	a.counts[a.index] = 0
}

func (a *Alerter) DisplayString() string {
	return a.alerts.DisplayAlerts()
}

func (a *Alerter) Flush() {
}

// Computes the average number of requests per second in the AlertDelay * Tick seconds
func (a *Alerter) computeAverage() float64 {
	nb := a.cfg.AlertDelay
	if a.firstPass {
		nb = a.index
	}
	totalDelay := float64(nb * a.cfg.Tick)
	sum := 0.0
	for i := 0; i < nb; i++ {
		sum += float64(a.counts[i])
	}
	return sum / totalDelay
}
