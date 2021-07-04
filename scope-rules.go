package main

import (
	"fmt"
	"math/rand"
)

var era = "公元"

func main() {
	year := 2021
	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, "年", month, "月", day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, "年", month, "月", day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, "年", month, "月", day)
	}
}
