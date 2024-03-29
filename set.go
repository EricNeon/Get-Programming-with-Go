package main

import (
	"fmt"
	"sort"
)

//按每10度进行分类,注意每次的结果由于是map，不讲究顺序
func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("预设值出现！")
	}
	fmt.Println(set)
	//按照切片方式排序输出
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)

	fmt.Println(unique)
}
