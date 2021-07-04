package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if count := rand.Intn(10); count < 4 {
		fmt.Println("太空大冒险")
	} else if count == 7 {
		fmt.Println("天问一号")
	} else {
		fmt.Println("神舟十二号")
	}
}
