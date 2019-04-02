package main

import (
	"log"
)

func main() {
	// 印出 hello world
	log.Println("standard: hello world")

	// 設定輸出的格式
	log.SetFlags(log.Ltime)
	log.Println("setflag:  hello world")

	// 設定前綴
	log.SetPrefix("[Bug] ")
	log.Println("prefix: hello world")

	// 印出資訊後，呼叫 exit 離開程式
	log.Fatal("Fatal: followed by a call to os.Exit(1)")
}
