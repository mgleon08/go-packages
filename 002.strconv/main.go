package main

import (
	"fmt"
	"log"
	"strconv"
)

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 字串轉數字
	i, err := strconv.Atoi("8")
	errHandler(err)
	fmt.Println("to int:", i)

	// 數字轉字串
	// Itoa is shorthand for FormatInt(int64(i), 10).
	fmt.Println("to string:", strconv.Itoa(8))

	// 字串轉 bool
	b, err := strconv.ParseBool("true")
	errHandler(err)
	fmt.Println("to bool:", b)

	// 字串轉 float 可指定 bitSize
	f, err := strconv.ParseFloat("3.1415", 64)
	errHandler(err)
	fmt.Println("to float:", f)

	// 字串轉 int 可指定 base & bitSize
	ii, err := strconv.ParseInt("-42", 10, 64)
	errHandler(err)
	fmt.Println("to int:", ii)

	// 字串轉 uint 可指定 base & bitSize
	u, err := strconv.ParseUint("42", 10, 64)
	errHandler(err)
	fmt.Println("to uint:", u)

}
