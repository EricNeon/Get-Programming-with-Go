package main

import (
	"fmt"
	"math"
)

func main() {
	third := 10.0 / 3

	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%09.2f\n", third)

	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64)
	fmt.Println(pi32)

	var count float64
	fmt.Printf("%f\n", count)
	fmt.Println(count)
	fmt.Printf("%v\n", count)
}
