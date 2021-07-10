package main

import "fmt"

type kelvin float64
type celsius float64

//定义一个开尔文转换为摄氏度的函数
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

//类似上面的函数，这里使用了方法的定义
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	var c, c2 celsius

	//这里使用函数返回
	c = kelvinToCelsius(k)

	//这里使用方法返回
	c2 = k.celsius()

	//两种返回应该是一样的结果
	fmt.Println(c)
	fmt.Println(c2)
}
