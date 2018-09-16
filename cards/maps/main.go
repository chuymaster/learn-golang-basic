package main

import (
	"fmt"
)

func main() {

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// fmt.Println(colors)

	// delete(colors, "white")

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {

	// Map is reference type.
	// Struct is value type
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
