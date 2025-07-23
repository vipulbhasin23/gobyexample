package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index of 3:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
