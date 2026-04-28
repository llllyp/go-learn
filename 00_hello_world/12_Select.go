package main

import (
	"fmt"
	"time"
)

/*
Go 的 select = 专门监听 Channel（通道）的「多路复用器」
同时等待多个 Channel 的读 / 写
哪个 Channel 先就绪，就执行哪个 case
所有 Channel 都阻塞 → 要么等待，要么走 default
Go 语言级别的「IO 多路复用」

select {
case <-ch1:
    // ch1 有数据了，执行这里
case ch2 <- value:
    // 往 ch2 写数据成功，执行这里
default:
    // 所有 channel 都阻塞，立刻走这里（不阻塞）
}

# 四大经典使用场景
场景 1：同时等待多个通道（谁先来执行谁）
```go
select {
case <-chA:
    fmt.Println("甲完成了")
case <-chB:
    fmt.Println("乙完成了")
}
```
场景 2：超时控制（超级常用！）
```go
select {
case <-ch:
    fmt.Println("收到数据")
case <-time.After(1 * time.Second):
    fmt.Println("超时了！") // 1 秒没收到就走这里
}
```
场景 3：非阻塞读写（不卡住程序）
```go
select {
case <-ch:
    fmt.Println("有数据")
default:
    fmt.Println("没数据，跳过不等待")
}
```
场景 4：优雅退出协程（关闭通知）
select {
case <-exitChan:
    fmt.Println("收到退出信号，停止工作")
    return
default:
    // 正常工作
}
*/

// 甲乙交替执行
func main() {
	chA := make(chan struct{})
	chB := make(chan struct{})

	// 甲
	go func() {
		for {
			select {
			case <- chA:
				fmt.Println("甲执行任务")
				time.Sleep(1 * time.Second)
				chB <- struct{}{}  // 往B通道写入一个数据, 想当于通知Bchannel
			}
		}
	}()

	// 乙
	go func() {
		for {
			select {
			case <- chB:
				fmt.Println("乙执行任务")
				time.Sleep(1 * time.Second)
				chA <- struct{}{}  // 往A通道写入一个数据, 想当于通知Achannel
			}
		}
	}()

	// 启动, 让甲先执行
	chA <- struct{}{}

	// 阻塞程序
	select{}
}