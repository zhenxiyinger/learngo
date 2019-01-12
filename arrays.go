package main

import "fmt"

func main() {
	arr := [...]int{2, 3, 4, 5, 6, 7}
	fmt.Println(arr)

	var grid [4][5]int
	fmt.Println(grid)

	for i, v := range arr {
		fmt.Println(i, v)
	}

	s := arr[1:4]
	fmt.Println(s)

}
