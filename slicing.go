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

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants)

	fmt.Println(gasGiants[0])

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice)

	陆地行星 := planets[:4]
	气态行星 := planets[4:6]
	冰态行星 := planets[6:]

	fmt.Println(陆地行星, 气态行星, 冰态行星)

	allPlants := planets[:]
	fmt.Println(allPlants)
}
