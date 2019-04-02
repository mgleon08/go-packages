package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字串轉換大小寫
	upperString, lowerString := "HELLO WORLD", "hello world"
	fmt.Println("ToUpper:", strings.ToUpper(lowerString))
	fmt.Println("ToLower:", strings.ToLower(upperString))

	// 指定特定分隔符號轉換成 slice
	splitString := "a,b,c"
	split := strings.Split(splitString, ",")
	fmt.Println("Split:", split)

	// 依照特定分隔符，合併成一個 string
	joinSlice := []string{"a", "b", "c"}
	fmt.Println("Join:", strings.Join(joinSlice, ","))

	// 字串1是否有包含字串2
	fmt.Println(strings.Contains("seafood", "foo"))

	// 字串替換，最後一個參數是指最多取代幾個，-1 就是全部都取代
	fmt.Println(strings.Replace("https://yahoo.com", "yahoo", "google", -1))
}
