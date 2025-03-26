package example

import (
	"demo/future"
	"fmt"
	"testing"
	"time"
)

func TestFuture(t *testing.T) {
	f := future.NewFuture()

	// 启动异步任务
	go future.AsyncTask(f)

	// 获取结果，设置超时时间为 3 秒
	result, err := f.Get(3 * time.Second)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
