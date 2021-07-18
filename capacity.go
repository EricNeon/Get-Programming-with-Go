package main

import "fmt"

func dump(lable string, slice []string) {
	fmt.Printf("%v:  长度 %v, 容量 %v %v\n", lable, len(slice), cap(slice), slice)
}
func main() {
	//初始化一个容量只有3的切片空间
	allplanets := make([]string, 0, 3)
	dump("allplanets", allplanets)
	//当添加的新值超过了原有切片容量时，自动扩充一倍的容量
	planets := []string{"水星", "金星", "地球", "火星", "木星", "土星", "天王星", "海王星"}
	for i := range planets {
		allplanets = append(allplanets, planets[i])
		dump("allplanets", allplanets)

	}

}
