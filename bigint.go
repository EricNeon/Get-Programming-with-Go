package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792)
	//lightSpeed := big.NewInt(30e30) //这样写会报错

	secondsPerDay := big.NewInt(86400)
	fmt.Println(lightSpeed, secondsPerDay)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println(distance)
}
