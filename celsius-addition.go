package main

import "fmt"

func main() {
	type celsius float64
	const degrees = 20
	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)

	//需要先转换为同一类型才能进行计算
	var temperature2 celsius = degrees
	var warmUp float64 = 33
	temperature2 += celsius(warmUp)
	fmt.Println(temperature2)
}
