package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var jsonString string
	jsonString = `{"name":"syhlion","age":5}`

	//把 json unmarshal 進去 struct
	u := &UserInfo{}
	err := json.Unmarshal([]byte(jsonString), u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("name:%s, age:%d\n", u.Name, u.Age)

}
