package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

//开尔文转摄氏度
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

//华氏度转摄氏度
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

//摄氏度转开尔文
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

//华氏度转开尔文
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

//开尔文转华氏度
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

//摄氏度转华氏度
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func main() {
	var k kelvin = 294.0
	var c celsius

	//这里使用方法返回
	c = k.celsius()

	//两种返回应该是一样的结果
	fmt.Printf("%v K 相当于%.2f摄氏度", k, c)

}
