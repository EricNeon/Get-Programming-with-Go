package main

import (
	"fmt"
	"math/rand"
)

var era = "公元"

func main() {
	for count := 0; count < 10; count++ {
		year := 2020 + rand.Intn(100)
		month := 2
		daysInMonth := 31
		leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
		switch month {
		case 2:
			if leap {
				fmt.Println("是闰年")
				daysInMonth = 29
			}
			daysInMonth = 28
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, "年", month, "月", day, "日")

	}
}
