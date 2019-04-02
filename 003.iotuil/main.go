package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %T\nValue: %v\nString: %s\n", content, content, content)

	r := strings.NewReader("hello")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}
