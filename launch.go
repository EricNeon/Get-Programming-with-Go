package main

import (
	"fmt"
)

func main() {
	launch := false

	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("准备用餐？:", launchText)

	var yesNo string
	if launch {
		yesNo = "是的！"
	} else {
		yesNo = "否！"
	}
	fmt.Println("准备用餐？：", yesNo)
}
