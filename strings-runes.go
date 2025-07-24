package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("Length:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count in string:", utf8.RuneCountInString(s))

	for idx, runeVal := range s {
		fmt.Printf("%#U starts at index: %d\n", runeVal, idx)
	}

	fmt.Println("\nDecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeVal, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeVal, i)
		w = width
	}
}
