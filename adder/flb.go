package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func stupid_fibonacci() func() int {
	a := 0
	b := 0
	c := 0
	return func() int {
		c += 1
		if c%2 == 1 {
			if a == 0 {
				a += 1
			} else {
				a = a + b
			}
			return a
		} else {
			if b == 0 {
				b += 1
			} else {
				b = a + b
			}
			return b
		}
	}
}

func fibonacci() intGenn {
	a, b := 0, 1
	return func() int {
		a, b = b, b+a
		return a
	}
}

type intGenn func() int

func (g intGenn) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	printFileContents(f)
}
