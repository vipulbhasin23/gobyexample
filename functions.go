package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("The sum is:", res)

	res = plusPlus(2, 3, 4)
	fmt.Println("The sum now is:", res)
}
