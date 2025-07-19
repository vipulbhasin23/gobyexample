package main

import "fmt"

func main() {
	var a [5]int

	fmt.Println("Array a:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("Length: ", len(a))
}
