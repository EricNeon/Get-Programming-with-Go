package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const myint = 4
	var count int

	for {
		count = rand.Intn(10)
		if count != myint {
			fmt.Println(count)
			fmt.Println("很遗憾，没猜对 *-^")
			time.Sleep(time.Second)
		} else {
			fmt.Println(count)
			fmt.Println("猜对了！就是", count)
			break
		}
	}

}
