package main

import "fmt"

func main() {
	const distance = 236000000000000000 // km
	const lightSpeed = 299792           // km/s
	const secondsPerDay = 86400         //s
	const daysPerYear = 365             //天

	const years = distance / lightSpeed / secondsPerDay / daysPerYear
	fmt.Println("大犬座矮星系距离地球大约", years, "光年。")
}
