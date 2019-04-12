package main

import (
	"fmt"
	"os"

	"github.com/hpcloud/tail"
)

// Reads all new lines and send them to the channel
func consume(file string, msgChan chan<- string) {
	stream, err := tail.TailFile(file, tail.Config{
		Follow:   true,
		Location: &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END},
	})
	if err != nil {
		fmt.Printf("Unable to open file \"%s\"", file)
		return
	}

	for line := range stream.Lines {
		msgChan <- line.Text
	}
}
