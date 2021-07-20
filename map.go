package main

import "fmt"

func main() {
	temperature := map[string]int{
		"地球": 15,
		"火星": -65,
	}

	temp := temperature["地球"]
	fmt.Printf("地球上的平均温度是：%v 摄氏度。\n", temp)
	temperature["地球"] = 16
	temperature["金星"] = 464

	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon)
	//推荐使用下面的判断语句，以避免空值输出
	if moon, found := temperature["Moon"]; found {
		fmt.Printf("月球上的平均温度是%v 摄氏度。\n", moon)
	} else {
		fmt.Println("月球在哪里？")
	}

}
