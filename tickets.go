package main

import (
	"fmt"
	"math/rand"
)

const distance = 62100000 //2020年10月13日火星到地球的距离 km
const secondsPerDay = 86400

func main() {
	var company, trip = "", ""

	fmt.Println("航班			 		飞行天数					飞行类型			价格（百万RMB）")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "SpaceX-9号"
		case 1:
			company = "神舟-18号"
		case 2:
			company = "问天-3号"
		}
		speed := rand.Intn(15) + 16                  //在15~30 km/s范围内选择
		duration := distance / speed / secondsPerDay //天数
		price := 20.0 + speed
		if rand.Intn(2) == 1 {
			trip = "往返"
			price = price * 2
		} else {
			trip = "单程"
		}
		fmt.Printf("%-20v %-20v %-20v $%-20v\n", company, duration, trip, price)

	}

}
