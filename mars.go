package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s

	const hoursPerDay, minutesPerHour, secondsPermin = 24, 60, 60
	var speed float32 = 11.20 // km/s
	var distance = 96300000   //km

	fmt.Println("从地球到达火星需要：", float32(distance)/speed/secondsPermin/minutesPerHour/hoursPerDay, "天")
	fmt.Println("从地球到达火星需要：", float32(distance)/speed/minutesPerHour, "分钟")
	fmt.Println("光从地球到达火星需要：", distance/lightSpeed, "秒")

}
