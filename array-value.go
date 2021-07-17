package main

import "fmt"

func main() {
	planets := [...]string{
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"天王星",
		"海王星",
	}

	planetsMarkII := planets

	planets[2] = "彗星"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}
