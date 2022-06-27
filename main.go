package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	eb.printGreeting()
	sb.printGreeting()
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (eb englishBot) printGreeting() {
	fmt.Println(eb.getGreeting())
}

func (sb spanishBot) printGreeting() {
	fmt.Println(sb.getGreeting())
}
