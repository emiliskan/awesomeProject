package main

import "fmt"

func main() {

	colors := map[string]string{
		"red": "#ff000",
	}

	printMap(colors)
	colors["yellow"] = "#yellow"
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
