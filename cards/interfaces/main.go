package main

import (
	"fmt"
)

type bot interface {
	// Any type that implements getGreeting() automatically inherits this interface.
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	pringGreeting(eb)
	pringGreeting(sb)
}

func pringGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello World"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}
