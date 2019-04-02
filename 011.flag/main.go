package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagTest string
	flag.StringVar(&flagTest, "flagTest", "default describe", "help information")

	flag.Parse()
	fmt.Printf("%v", flagTest)
}

// go run main.go --flagTest=hi
