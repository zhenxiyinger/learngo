package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a % b
	default:
		panic("错误符号")
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)", opName, a, b)
	return op(a, b)
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	a, b := div(13, 3)
	fmt.Println(
		a,
		b,
		apply(
			func(a, b int) int {
				return int(math.Pow(float64(a), float64(b)))
			}, 3, 4),
		apply(pow, 3, 4),

		sum(1, 2, 3, 4, 5),
	)

	c, d := 10, 20
	c, d = swap(c, d)
	println(c, d)
}

func swap(a, b int) (int, int) {
	return b, a
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
