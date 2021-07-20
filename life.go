package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

//Universe做一个二维网格
type Universe [][]bool

//NewUniverse 将返回一个空白的世界
func NewUniverse() Universe {
	u := make(Universe, height, width)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

//Seed 用于在世界中随机产生一些细胞
func (u Universe) Seed() {
	for i := 0; i < (width * height / 15); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

//Set 用于设置指定细胞的状态
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

//Alive会报告指定的细胞是否存活，如果给定的坐标不在世界范围内，则实行回绕
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

//Neighbors用于统计邻近细胞的存活数量
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

//Next会返回指定细胞下一个世代的状态
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	//	return n == 3 || n == 2 && u.Alive(x, y)
	return n == 5 || n == 1 || n == 2 || n == 3 || n == 4 && u.Alive(x, y)
}

//String会以字符串形式返回整个世界
func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

//show清空屏幕并打印整个世界
func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}

//Step会将当前世界a的状态更新至下一世代并将其存储在世界b中
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 30; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 3)
		a, b = b, a //Swap Universe
	}
}
