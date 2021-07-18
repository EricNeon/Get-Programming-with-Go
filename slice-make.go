package main

import "fmt"

func dump(lable string, slice []string) {
	fmt.Printf("%v:  长度 %v, 容量 %v %v\n", lable, len(slice), cap(slice), slice)
}
func main() {

	planets := make([]string, 0, 10)
	dump("planets", planets)

	planets = append(planets, "水星", "金星", "地球", "火星", "木星", "土星", "天王星", "海王星")
	dump("planets", planets)

}
