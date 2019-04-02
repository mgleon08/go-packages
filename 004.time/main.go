package main

import (
	"fmt"
	"time"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

func main() {
	fmt.Println(Nanosecond, Microsecond, Millisecond, Second, Minute, Hour)

	t := time.Now()
	fmt.Println(t.Unix())     // seconds since 1970
	fmt.Println(t.UnixNano()) // nanoseconds since 1970

	// 解析時間，並轉成分鐘
	m, _ := time.ParseDuration("1h30m")
	fmt.Printf("The movie is %.0f minutes long.\n", m.Minutes())

	// format 格式化後的時間
	// 後面格式是裡面的設定，照 1 ~ 7 去記 01/02 03:04:05PM '06 -0700
	// https://golang.org/src/time/format.go
	fmt.Println("datetime: ", time.Now().Format("2006-01-02T15:04:05Z07:00"))

	// 停多少時間後再執行
	// AfterFunc waits for the duration to elapse and then calls f
	// in its own goroutine. It returns a Timer that can be used to
	// cancel the call using its Stop method.
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("hello world")
	})

	// 這裡 sleep 主要是因為上面的 AfterFunc 會跑 goroutine，沒有的話 main thread 就會直接結束
	time.Sleep(4 * time.Second)
	fmt.Println("end")
}
