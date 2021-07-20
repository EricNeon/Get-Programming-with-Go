package main

import "fmt"

func main() {
	planets := map[string]string{
		"地球": "Sector zz9",
		"火星": "Sector zz9",
	}

	planetsMarkII := planets
	planets["地球"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	delete(planets, "地球")
	fmt.Println(planetsMarkII)
}
