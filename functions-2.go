package main

import "fmt"

func vals() (int, int) {
	return 5, 15
}

func main() {
	v1, v2 := vals()
	fmt.Println("The two return values:", v1, v2)

	_, v2 = vals()
	fmt.Println("The second return value is:", v2)
}
