package main

import "fmt"

func main() {
	fmt.Println("2100年你能起飞吗？（是闰年吗？）")

	var year = 2100

	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("是闰年，请三思而后行！")
	} else {
		fmt.Println("不是闰年！")
	}
}
