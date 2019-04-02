package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	re := regexp.MustCompile("(gopher){2}")
	fmt.Println(re.MatchString("gopher"))
	fmt.Println(re.MatchString("gophergopher"))
	fmt.Println(re.MatchString("gophergophergopher"))

	re = regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))

	re, _ = regexp.Compile(`http`)
	reString := re.ReplaceAllString("http://google.com", `https`)
	fmt.Printf("%q\n", reString)

	re, _ = regexp.Compile(`123`)
	if len(os.Args) > 1 {
		fmt.Printf("%q\n", re.FindStringSubmatch(os.Args[1]))
	}
}
