package main

import "fmt"

func main() {
	fmt.Println("祝你平安！\n祝你平安。")
	fmt.Println(`原始字符串可以用 \n 符号切分为多行。`)
	fmt.Println("原始字符串可以用", "\\n", "符号切分为多行。")
	fmt.Println("原始字符串可以用", `\n`, "符号切分为多行。")

	fmt.Println(`
		这里会保留换行格式
		不会破坏原来的结构。
		`)

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	//将会打印出code point的值

	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)
	//将会打印出字符

	//打印出星号*、笑脸和重音符号e的代码点
	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)
	smile := '☺'
	fmt.Printf("%c %[1]v\n", smile)
	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)

}
