package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

//sensor函数类型
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5
	sensor := calibrate(fakeSensor, offset)
	for i := 0; i < 10; i++ {
		fmt.Println(sensor())
	}

	//下面是显示引用值时发生的变化
	/*	var k kelvin = 294.0
		sensor = func() kelvin {
			return k
		}
		fmt.Println(sensor())
		k++
		fmt.Println(sensor())*/

}
