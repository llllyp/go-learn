package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
func someFunc() {
	// do something
}
// 常见错误和陷阱
func main() {

	// 错误1: 没有等待 goroutine
	go someFunc() //goroutine 还没执行完, main就结束了

	// 正确思路
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		someFunc()
	}()
	wg.Wait()

	// 错误2: 闭包捕获循环变量
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i) // 输出可能是3
		}()
	}
	// 正确
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Println(id) // 输出正确的值
		}(i)
	}

	// 错误3: channel没有关闭

	// 错误4: 在 goroutine 中不使用 recover
	go func() {
		panic("oops") //会导致整个程序崩溃
	}()

	// 正确
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("捕获的panic:", r)
			}
		}()
		// 可能 panic 的代码
	}()

	// --------------- 性能建议
	// 1. 避免频繁创建 goroutine
	for i := 0; i < 3; i++ { // ❌ 在循环中重复创建
		go func(it int) {
			// do something
		}(i)
	}
	// 正确 使用 worker pool
func workerPool(items []Item, numWorkers int) {
		jobs := make(chan Item, len(items))

		// 启动 worker
		for w := 0; w < numWorkers; w++ {
			go func ()  {
				for item := range jobs {
					process(item)
				}
			}()
		}

		// 发送任务
		for _, item := range items {
			jobs <- item
		}
		close(jobs)
		//等待完成
	}

	// 2. 使用缓冲 channel 提高性能
	ch := make(chan int, 100)

	// 3, 合理使用 Context 超时
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	fucnc process(item int) {
	
	}
}


