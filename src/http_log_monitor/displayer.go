package main

import "fmt"

type Displayer interface {
	Display(line string)
}

type ConsoleDisplayer struct {
}

func (d *ConsoleDisplayer) Display(line string) {
	if len(line) > 0 {
		fmt.Println(line)
	}
}
