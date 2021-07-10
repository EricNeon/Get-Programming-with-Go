package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown := 10
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str)

	countdown2 := 9
	str2 := fmt.Sprintf("Launch in T minus %v seconds.", countdown2)
	fmt.Println(str2)
}
