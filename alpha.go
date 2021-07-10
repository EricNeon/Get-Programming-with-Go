package main

import (
	"fmt"
)

//计算飞到半人马座阿尔法星所需的天数
func main() {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12

	fmt.Println("阿尔法星距离地球", distance, "km远。")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("以光速旅行至阿尔法星需要", days, "天时间。")

	var Ndistance = 24e18 //超出uint64范围，可以用float64类型
	fmt.Println("距离仙女座有", Ndistance, "km远。")
}
