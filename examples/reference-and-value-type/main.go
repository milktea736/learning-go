package main

import "fmt"

func main() {
	colors := map[string]int{
		"red":   1,
		"blue":  2,
		"green": 3,
	}

	updateMap(colors, "red", 10)
	fmt.Println(colors)

	myColors := myMap{
		"red":   1,
		"blue":  2,
		"green": 3,
	}
	myColors.updateMap("red", 10)
	fmt.Println(myColors)
}

type myMap map[string]int

func (c myMap) updateMap(color string, value int) {
	c[color] = value
}

func updateMap(c map[string]int, color string, value int) {
	c[color] = value
}
