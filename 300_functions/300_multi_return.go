package main

import "fmt"

func cutInHalf(text string) (string, string) {
	return text[:len(text)/2], text[len(text)/2:]
}

func main() {
	text := "alamakota"

	a, b := cutInHalf(text)
	fmt.Printf("A: %s, B: %s", a, b)
}