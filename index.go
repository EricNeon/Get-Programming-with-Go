package main

import "fmt"

func main() {
	peace := "shalom.hello world,welcome to ..."
	c := peace[3:15]
	for i, value := range c {
		fmt.Printf("%v %c\n", i, value) //i作为字符所在的索引位置，value代表字符本身
	}
}
