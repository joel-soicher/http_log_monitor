package main

import "strconv"

// Defines the base entity that will be used for stats, alerts, ...
type Checker interface {
	AddRequest(req *Request)
	Compute()
	DisplayString() string
	Flush()
}

// Checker that will count the number of invalid lines in the file
type InvalidChecker struct {
	count int64
}

func (c *InvalidChecker) AddRequest(req *Request) {
	if req == nil {
		return
	}
	if !req.Valid {
		c.count++
	}
}

func (c *InvalidChecker) Compute() {
}

func (c *InvalidChecker) DisplayString() string {
	return "Invalid line: " + strconv.FormatInt(c.count, 10)
}

func (c *InvalidChecker) Flush() {
	c.count = 0
}

// Checker that will count the number of request with 200 status
type OkChecker struct {
	count int64
}

func (c *OkChecker) AddRequest(req *Request) {
	if req == nil {
		return
	}
	if req.Status == 200 {
		c.count++
	}
}

func (c *OkChecker) Compute() {
}

func (c *OkChecker) DisplayString() string {
	return "200 Status: " + strconv.FormatInt(c.count, 10)
}

func (c *OkChecker) Flush() {
	c.count = 0
}
