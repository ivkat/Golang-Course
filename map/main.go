package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ddffg0",
		"white": "#ffffff",
	}

	//or...
	//colors := make(map[string]string)

	//or...
	//var colors map[string]string

	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
