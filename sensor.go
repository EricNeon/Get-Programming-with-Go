package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

//温度传感器读取150k-300k区间内的气温
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func main() {
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}
