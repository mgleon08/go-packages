package main

import (
	"errors"
	"fmt"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("b can not be 0")
	}
	return a / b, nil
}

func main() {
	fmt.Println(Division(60, 2))
}
