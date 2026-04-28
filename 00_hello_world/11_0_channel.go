package main

import "fmt"

/*
并发核心单元通过它就可以发送或者接收数据进行通讯
goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
定义channel变量
    make(chan Type)  //等价于make(chan Type, 0)
    make(chan Type, capacity)
当复制一个channel或用于函数参数传递时，只是拷贝了一个channel引用，
因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

当 参数capacity= 0 时，channel 是无缓冲阻塞读写的；当capacity > 0 时，channel 有缓冲、是非阻塞的，直到写满 capacity个元素才阻塞写入。

    channel <- value      //发送value到channel
    <-channel             //接收并将其丢弃
    x := <-channel        //从channel中接收数据，并赋值给x
    x, ok := <-channel    //功能同上，同时检查通道是否已关闭或者是否为空

关闭channel后，无法向channel 再发送数据(引发 panic 错误后导致接收立即返回零值)；

关闭channel后，可以继续从channel接收数据；

对于nil channel，无论收发都会被阻塞。

可以使用 range 来迭代不断操作channe

单向channel
	var ch1 chan int       // ch1是一个正常的channel，是双向的
	var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据
	var ch3 <-chan int     // ch3是单向channel，只用于读int数据
*/

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子go程结束")
		fmt.Println("子go程正在运行")

		c <- 666 //将 666 发送到C
	}()

	num, ok := <- c  // 从 C 中接收数据, 并赋值给 num
	fmt.Println("num = ", num)
	if ok {
		fmt.Println("通道 c 未关闭")
		close(c) // 关闭通道 c
	} else {
		fmt.Println("通道 c 已关闭")
	}
	fmt.Println("main go程结束")

}