package main

import (
	"fmt"
	"strconv"
)

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
	return "Invalid lines: " + strconv.FormatInt(c.count, 10)
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

// Checker that will compute min, max and average size of requests
type SizeChecker struct {
	count int64
	total uint64
	min   uint64
	max   uint64
}

func NewSizeChecker() *SizeChecker {
	return &SizeChecker{
		min: ^uint64(0),
	}
}

func (c *SizeChecker) AddRequest(req *Request) {
	if req == nil {
		return
	}
	c.count++
	b := req.Bytes
	c.total += b
	if b <= c.min {
		c.min = b
	}
	if b >= c.max {
		c.max = b
	}
}

func (c *SizeChecker) Compute() {
}

func (c *SizeChecker) DisplayString() string {
	if c.count == 0 {
		return ""
	}
	avg := float64(c.total) / float64(c.count)
	return fmt.Sprintf("Average request size:  %.2f (min: %v, max: %v)", avg, c.min, c.max)
}

func (c *SizeChecker) Flush() {
	c.count = 0
	c.total = 0
	c.min = 0
	c.max = 0
}
