package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(n, "監控退出，停止了...")
			return
		default:
			fmt.Printf("goroutine %v 監控中...\n", n)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	n := 3
	for i := 0; i < n; i++ {
		go process(ctx, i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("通知監控停止")
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("所有監控完成")
}
