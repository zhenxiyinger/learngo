package main

import "fmt"

func searchMastLong(str string) int {
	last := make(map[rune]int)
	start := 0
	max := 0

	for i, ch := range []rune(str) {
		if lastI, ok := last[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > max {
			max = i - start + 1
		}

		last[ch] = i
	}
	return max
}

func main() {
	fmt.Println([]byte("abcabcacbcc"))
	//fmt.Println(searchMastLong("abcabcacbcc"))
}
