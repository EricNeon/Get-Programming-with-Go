package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println("倒计时: ", count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("发射！")
}
