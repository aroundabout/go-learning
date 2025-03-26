package future

import (
	"fmt"
	"time"
)

// Future 结构体模拟 Java 的 Future 特性
type Future struct {
	resultChan chan interface{}
	errChan    chan error
}

// NewFuture 创建一个新的 Future 实例
func NewFuture() *Future {
	return &Future{
		resultChan: make(chan interface{}),
		errChan:    make(chan error),
	}
}

// SetResult 设置 Future 的结果
func (f *Future) SetResult(result interface{}) {
	f.resultChan <- result
	close(f.resultChan)
}

// SetError 设置 Future 的错误信息
func (f *Future) SetError(err error) {
	f.errChan <- err
	close(f.errChan)
}

// Get 获取 Future 的结果，支持超时
func (f *Future) Get(timeout time.Duration) (interface{}, error) {
	select {
	case result := <-f.resultChan:
		return result, nil
	case err := <-f.errChan:
		return nil, err
	case <-time.After(timeout):
		return nil, fmt.Errorf("timeout after %v", timeout)
	}
}

// AsyncTask 模拟一个异步任务
func AsyncTask(f *Future) {
	time.Sleep(2 * time.Second)
	// 模拟任务成功
	f.SetResult(42)
	// 模拟任务失败
	// f.SetError(fmt.Errorf("task failed"))
}
