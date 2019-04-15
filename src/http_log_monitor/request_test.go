package main

import (
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	const s = "127.0.0.1 - james [09/May/2018:16:00:39 +0000] \"GET /report HTTP/1.0\" 200 123"
	r := Parse(s)

	if r.Valid == false ||
		r.HostName != "127.0.0.1" ||
		r.User != "-" ||
		r.AuthUser != "james" ||
		!r.Date.Equal(time.Date(2018, time.May, 9, 16, 0, 39, 0, time.FixedZone("+0000", 0))) ||
		r.Method != "GET" ||
		r.Resource != "/report" ||
		r.Protocol != "HTTP/1.0" ||
		r.Status != 200 ||
		r.Bytes != 123 {
		t.Error()
	}
}

func TestParserInvalid(t *testing.T) {
	const s = "invalid string"
	r := Parse(s)
	if r.Valid {
		t.Error()
	}
}
