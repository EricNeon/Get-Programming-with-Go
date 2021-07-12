package main

import "fmt"

func main() {
	func() {
		fmt.Println("这是匿名函数！")
	}() //最后的()代表立即执行这个匿名函数
}
