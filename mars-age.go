package main

import (
	"fmt"
	"math"
)

func main() {
	age := 43
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Printf("在火星上，我的年龄是 %05.2f 岁。\n", marsAge)

	fmt.Println(int(earthDays))

	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("转换为:", v8)
	}
}
