package main

import (
	"testing"
)

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //呼叫該函式停止壓力測試的時間計數

	//做一些初始化的工作，例如讀取檔案資料，資料庫連線之類別的,
	//這樣這些時間不影響我們測試函式本身的效能

	b.StartTimer() //重新開始時間
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
