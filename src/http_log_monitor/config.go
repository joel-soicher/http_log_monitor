package main

type Config struct {
	Tick        int // Time interval between each display
	AlertDelay  int // Number of intervals in which alerts must be checked
	MaxReq      int // Max number of requests triggering an alert
	MaxSections int // Max number of most hit sections
}
