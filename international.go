package main

import "fmt"

func main() {
	message := "Chinese premier holds dialogue with British business leaders"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}

		fmt.Printf("%c", c)
	}
}
