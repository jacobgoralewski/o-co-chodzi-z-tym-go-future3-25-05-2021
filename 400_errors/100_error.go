package main

import (
	"errors"
	"fmt"
	"log"
)

func div(a, b float64) (float64, error) {
	if b == 0. {
		return 0., errors.New("division by zero")
	}

	return a/b, nil
}

func main() {
	res, err := div(2., 1.)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println(res)
}
