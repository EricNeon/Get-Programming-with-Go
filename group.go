package main

import (
	"fmt"
	"math"
)

//按每10度进行分类,注意每次的结果由于是map，不讲究顺序
func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64) //值为一个float64的切片

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10 //向下取整
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%v:%v\n", g, temperatures)
	}
}
