package main

import (
	"fmt"
	"maps"
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

	fmt.Println("Number of key/value pairs:", len(m))

	delete(m, "k1")
	fmt.Println("Delete key/value pair for k1:", m)

	_, prs := m["k1"]
	fmt.Println("Optional second return value for k1:", prs)

	_, prs = m["k2"]
	fmt.Println("Optional second return value for k2:", prs)

	clear(m)
	fmt.Println("Clear map using the built-in clear:", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map literal:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
