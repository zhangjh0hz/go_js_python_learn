package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "YES我爱慕课网"
	var sb []string = strings.Fields(s)
	for _, ss := range sb {
		fmt.Println(ss)
	}
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}

	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

}
