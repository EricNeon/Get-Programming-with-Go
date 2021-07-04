package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("你发现自己在一个光线昏暗的洞穴里。")

	var command = "请从这里出去！"
	var exit = strings.Contains(command, "出去")

	fmt.Println("你可以离开洞穴吗？", exit)
}
