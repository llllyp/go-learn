package main

import (
	"fmt"
	"sync"
)

/*
单向 channel = 限制通道只能读 或 只能写
类型			    写法			能力
双向通道（正常）	 chan int		 可读、可写
只写通道			chan<- int		只能往里发，不能读
只读通道			<-chan int		只能往外读，不能发

双向通道 可以 自动转为 单向通道
ch := make(chan int)   // 双向

go producer(ch)        // 自动转成 只写 chan<-
go consumer(ch)        // 自动转成 只读 <-chan

*/
func main() {
	const totalTickets = 10 // 总票数

	// 1. 创建双向通道（会自动转为单向）
	ticketChan := make(chan struct{}, totalTickets)

	// 2. 等待组：等所有售票窗口结束
	var wg sync.WaitGroup

	// 3. 余票计数 + 锁
	var leftTickets = totalTickets
	var mu sync.Mutex // 余票锁

	// ======================
	// 生产者（放票）：只写通道 chan<-
	// ======================
	go producer(ticketChan, totalTickets)

	// ======================
	// 启动 4 个售票窗口
	// ======================
	wg.Add(4)
	go consumer("窗口1", ticketChan, &wg, &leftTickets, &mu)
	go consumer("窗口2", ticketChan, &wg, &leftTickets, &mu)
	go consumer("窗口3", ticketChan, &wg, &leftTickets, &mu)
	go consumer("窗口4", ticketChan, &wg, &leftTickets, &mu)

	// 等待所有售票结束
	wg.Wait()
	close(ticketChan)
	fmt.Println("\n🎉 所有票卖完！程序正常退出")
}

// ======================
// 生产者：只写通道，禁止读！安全！
// ======================
func producer(ch chan<- struct{}, total int) {
	for i := 0; i < total; i++ {
		ch <- struct{}{} // 放一张票
	}
	// 放完关闭通道 → 告诉消费者：没票了
	close(ch)
}

// ======================
// 消费者：只读通道，禁止写！安全！
// ======================
func consumer(
	name string,
	ch <-chan struct{},
	wg *sync.WaitGroup,
	left *int,
	mu *sync.Mutex,
) {
	defer wg.Done()

	// 遍历只读通道，有票就卖
	for range ch {
		// 卖票
		fmt.Printf("\n✅ %s 卖出一张票\n", name)

		// 修改余票（加锁保证安全）
		mu.Lock()
		*left--
		fmt.Printf("📊 剩余票数：%d\n", *left)
		mu.Unlock()
	}

	// 遍历结束 = 通道关闭 + 无数据
	fmt.Printf("\n🏠 %s 所有票卖完，下班！\n", name)
}