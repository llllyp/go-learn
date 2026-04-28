package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 总票数
	total := 10

	ticketChan := make(chan int, total)
	var wg sync.WaitGroup
	var mu sync.Mutex // 用来保护余票计数

	// 余票变量（我们自己维护）
	left := total

	// ========== 1. 生产者：放票 ==========
	go func() {
		for i := 0; i < total; i++ {
			ticketChan <- 1
		}
		close(ticketChan) // 放完就关闭，不影响卖票
	}()

	// ========== 2. 卖票逻辑 ==========
	sell := func(name string) {
		defer wg.Done()

		// 从通道取票
		for range ticketChan {
			// 卖票
			fmt.Printf("\n✅ %s 卖出一张票！\n", name)

			// 修改余票（加锁！）
			mu.Lock()
			left--
			fmt.Printf("📊 余票：%d 张\n", left)
			mu.Unlock()

			// 模拟卖票慢
			time.Sleep(1 * time.Second)
		}

		fmt.Printf("\n🏠 %s 下班了\n", name)
	}

	// ========== 3. 开启 4 个窗口 ==========
	wg.Add(4)
	go sell("窗口1")
	go sell("窗口2")
	go sell("窗口3")
	go sell("窗口4")

	// ========== 4. 等待所有窗口卖完 ==========
	wg.Wait()

	fmt.Println("\n========================")
	fmt.Println("🎉 所有票全部卖完！")
	fmt.Println("========================")
}