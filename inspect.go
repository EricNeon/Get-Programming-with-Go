package main

import "fmt"

func main() {
	year := 2021
	fmt.Printf("Type %T for %[1]v\n", year)

	a := "ok"
	fmt.Printf("Type %T for %[1]v\n", a)

	c := 3.14159
	fmt.Printf("Type %T for %[1]v\n", c)

	n := false
	fmt.Printf("Type %T for %[1]v\n", n)
}
