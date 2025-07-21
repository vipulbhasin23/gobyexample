package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Map:", m)

	v1 := m["k1"]
	fmt.Println("Value for k1:", v1)

	v3 := m["k3"]
	fmt.Println("Value for k3(not present):", v3)
}
