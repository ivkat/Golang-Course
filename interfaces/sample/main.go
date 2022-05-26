package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// We can't overload like the commonly used OOP languages do
// For example, printGreeting(sb spanishBot){} would fail to compile here
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// No variable name since we're not doing anything. Just adding behaviour to the type this time
func ( /*eb*/ englishBot) getGreeting() string {
	return "Hello there"
}

func ( /*sb*/ spanishBot) getGreeting() string {
	return "Hola"
}
