package main

import "fmt"

func main() {
	var haveTorch = true
	var litTorch = false
	if !haveTorch || !litTorch {
		fmt.Println("什么也没发现!")
	}
}
