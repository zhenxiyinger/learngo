package main

import "fmt"

func main() {
	s := "Yes我爱慕课网!"
	fmt.Println(len(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	for i, r := range []rune(s) {
		fmt.Printf("(%d %c) ", i, r)
	}
	fmt.Println()
}
