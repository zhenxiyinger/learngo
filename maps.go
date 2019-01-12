package main

import (
	"fmt"
)

func main() {
	my := map[string]string{
		"name": "tank",
		"age":  "31",
		"sex":  "男",
	}

	m2 := make(map[string]int)

	for k, v := range my {
		fmt.Println(k, v)
	}

	fmt.Println(my, m2)

	name := my["name"]
	fmt.Println(name)

	if age, ok := my["agee"]; ok {
		fmt.Println(age)
	}

	delete(my, "age")
	fmt.Println(my)

}
