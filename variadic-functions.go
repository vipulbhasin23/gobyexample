package main

import "fmt"

func sums(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("Sum is:", total)
}

func main() {
	sums(1)
	sums(2, 3, 4)

	nums := []int{1, 5, 6}
	sums(nums...)
}
