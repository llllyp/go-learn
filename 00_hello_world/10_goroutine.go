package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
协程
goroutine是Go语言并行设计的核心，有人称之为go程。 Goroutine从量级上看很像协程，它比线程更小，
十几个goroutine可能体现在底层就是五六个线程，Go语言内部实现了这些goroutine之间的内存共享。
执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。
也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。



*/

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 创建一个 goroutine, 启动另外一个任务
	// go newTask()
	// i :=0
	// // main goroutine 循环打印
	// for {
	// 	i++
	// 	fmt.Printf("main goroutine: i = %d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }

	go func(){
		defer fmt.Println("A.defer")

		func () {
			defer fmt.Println("B.defer")
			runtime.Goexit()// 终止当前 goroutine 
			fmt.Println("B.exit") //不会执行
		}()
		fmt.Println("A.exit") // 不会执行
	}()

	for{
		} // 死循环, 目的不让 主 goroutine 结束
}