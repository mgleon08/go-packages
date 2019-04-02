package main

import (
	"fmt"
)

func ExampleDivision() {
	i, _ := Division(6, 1)
	fmt.Print(i)
	// Output: 6
}

func ExampleDivision2() {
	i, _ := Division(6, 2)
	x, _ := Division(12, 3)
	fmt.Println(i)
	fmt.Println(x)
	// Output:
	// 3
	// 4
}
