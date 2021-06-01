package main

import "fmt"

func main() {
	var sub = func(a, b int) int {
		return a - b
	}

	res := sub(5, 3)
	fmt.Println(res)
}
