package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 4
	fmt.Println("i:", i)

	zeroval(i)
	fmt.Println("After zeroval(i):", i)

	zeroptr(&i)
	fmt.Println("After zeroptr(i):", i)

	fmt.Println("&i:", &i)
}
