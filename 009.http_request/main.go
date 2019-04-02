package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://mgleon08.github.io/")
	if err != nil {
		log.Fatal(err)
	}
	// res.Body 讀取完，記得 Close，不然會有 memory leak 等相關問題
	defer res.Body.Close()
	// 透過 ReadAll 讀出資訊，但出來的資訊會是 []byte
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// %s, 將 []byte 轉換成 string
	fmt.Printf("%s", html)
	// 也可以這樣寫
	// fmt.Printf("%v", string(html))
}
