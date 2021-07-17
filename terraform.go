package main

import "fmt"

func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "新" + planets[i]
	}
}
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

	terraform(planets)
	fmt.Println(planets)
}
