package main

import (
	"fmt"
	"math/rand"
)

var era = "公元"

func main() {
	year := 2020
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, "年", month, "月", day, "日")

}
