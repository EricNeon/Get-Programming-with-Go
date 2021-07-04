package main

import "fmt"

func main() {
	var command = "进去"
	if command == "往东方走" {
		fmt.Println("你继续往山上走。")
	} else if command == "进去" {
		fmt.Println("进入洞穴，你将在那里度过你的余生。")
	} else {
		fmt.Println("方向不明。")
	}
}
