package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
		if rand.Intn(100) < 50 {
			break
		}
	}
	if count == 0 {
		fmt.Println("点火！")
		time.Sleep(2 * time.Second)
		fmt.Println("起飞！")
	} else {
		fmt.Println("发现故障，任务失败，取消发射！")
	}
}
