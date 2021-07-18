package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	twoWorlds := terraform("旧", "金星", "火星")
	fmt.Println(twoWorlds)

	planets := []string{"金星", "火星", "木星"}
	newPlanets := terraform("新", planets...)
	fmt.Println(newPlanets)
}
