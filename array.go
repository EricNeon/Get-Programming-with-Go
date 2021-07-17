package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "水星Mercury"
	planets[1] = "金星Venus"
	planets[2] = "地球Earth"

	earth := planets[2]
	fmt.Println("第三颗行星是:", earth)

	fmt.Println("太阳系共有", len(planets), "颗行星。")
	fmt.Println(planets[3] == "") //判断是否为空
	//打印第一个值
	fmt.Println("第一颗行星是:", planets[0])
}
