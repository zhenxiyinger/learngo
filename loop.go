package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int) string {
	res := ""
	for ; n > 0; n /= 2 {
		res = strconv.Itoa(n%2) + res
	}
	return res
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		converToBin(5),
		converToBin(10),
		converToBin(2),
	)

	printFile("abc.txt")

}
