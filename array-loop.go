package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i := 0; i < len(dwarfs); i++ {
		dwarfs := dwarfs[i]
		fmt.Printf("第%v颗星是:%s\n", i+1, dwarfs)
	}

}
