package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch count := rand.Intn(10); count {
	case 0:
		fmt.Println("太空大冒险")
	case 1:
		fmt.Println("天问一号")
	case 2:
		fmt.Println("神舟十二号")
	default:
		fmt.Println("出发！")
	}
}
