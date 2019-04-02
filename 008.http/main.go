package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	u := &UserInfo{
		Name: "leon",
		Age:  18,
	}
	// 將 struct 轉 json
	b, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}
	// 設定回傳 json 格式
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	// 設定 200
	w.WriteHeader(http.StatusOK)
	// 將 json 寫入
	w.Write(b)
}

func main() {
	http.HandleFunc("/api/query", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
