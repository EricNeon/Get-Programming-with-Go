package main

import "fmt"

func main() {
	message := "L fdph,L vdz,L frqtxhuhg"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c < 'a' {
				c = c + 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c < 'a' {
				c = c + 26

			}
		}

		fmt.Printf("%c", c)
	}
}
