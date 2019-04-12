package main

import (
	"fmt"
	"time"
)

// Request in standard httpd format + additional fields
type Request struct {
	Valid    bool
	HostName string
	User     string
	AuthUser string
	Date     time.Time
	Method   string
	Resource string
	Protocol string
	Status   uint16
	Bytes    uint64
}

// Parses a line and fills a request
func Parse(line string) *Request {
	request := &Request{}
	var tmpDate, tmpTime, req string
	var err error
	if _, err = fmt.Sscanf(line,
		"%s %s %s [%s %5s] %q %d %d",
		&request.HostName,
		&request.User,
		&request.AuthUser,
		&tmpDate,
		&tmpTime,
		&req,
		&request.Status,
		&request.Bytes); err != nil {
		return request
	}
	const dateFormat = "02/Jan/2006:15:04:05 -0700"
	if request.Date, err = time.Parse(dateFormat, tmpDate+" "+tmpTime); err != nil {
		return request
	}
	if _, err = fmt.Sscanf(req, "%s %s %s",
		&request.Method,
		&request.Resource,
		&request.Protocol); err != nil {
		return request
	}
	request.Valid = true
	return request
}
