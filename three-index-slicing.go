package main

import "fmt"

func dump(lable string, slice []string) {
	fmt.Printf("%v:  长度 %v, 容量 %v %v\n", lable, len(slice), cap(slice), slice)
}
func main() {
	planets := []string{
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"天王星",
		"海王星"}
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")
	//	fmt.Println(planets)
	dump("planets", planets)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)

	fmt.Printf("-------------------------\n")
	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	//	fmt.Println(planets)
	dump("planets", planets)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)

}
