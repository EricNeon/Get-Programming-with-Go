package main

import "fmt"

func main() {
	fmt.Println("这里有一块指示牌，并有一条通往东方的路。")

	var command = "进入"

	switch command {
	case "向东方走":
		fmt.Println("你继续往山上走。")
	case "进洞", "进入":
		fmt.Println("你发现自己在一个昏暗的洞穴里。")
	case "读取信号":
		fmt.Println("标志上写着'禁止未成年人'。")
	default:
		fmt.Println("方向不明")
	}
}
