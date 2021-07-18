package main

import (
	"fmt"
	"sort"
)

func main() {
	message := []string{
		"Join", "UP", "ok", "Earth", "Moon", "Left", "Good", "Yes", "zoom", "wahaha",
	}
	sort.StringSlice(message).Sort()
	fmt.Println(message)
	//下面是简便的写法
	sort.Strings(message)
	fmt.Println(message)
}
