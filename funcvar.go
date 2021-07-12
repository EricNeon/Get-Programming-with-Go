package main

import "fmt"

func main() {
	f := func(message string) {
		fmt.Println(message)
	}

	f("出发去参加化妆舞会！")
}
