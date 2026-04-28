package main

import (
	"fmt"
)

/*
defer 语句 被用于预定一个函数的调用. 可以把这类被defer语句调用的函数称为延迟函数
作用:
	- 释放占用的资源
	- 捕捉处理异常
	- 输出日志

如果一个函数中有多个 defer语句, 会以 LIFO(后进先出) 的顺序执行

recover错误拦截
运行时 panic 异常一旦被引发就会导致程序崩溃
go语言提供了专用于"拦截"运行时 panic 的内建函数"recover". 它可以使当前的程序从运行时 panic 状态中恢复并重新获得流程控制权

recover 只有在 defer 调用的函数中有效

*/
func main() {
	testLifo() // 输出顺序 3 2 1

	testRecover(10) // 输出: 数组下标越界
	// 产生错误后,程序继续执行
	fmt.Println("程序继续执行......")
}

func testLifo() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func testRecover(i int) {
	// 定义 10 个元素的数组
	var arr [10]int
	// 错误拦截要在生产错误前设置
	defer func() {
		// 设置 recover 拦截错误信息
		err := recover()
		// 产生 panic 异常  打印错误信息
		if err != nil {
			fmt.Println("错误拦截:", err)
		}
	}()
	// 根据函数参数为数组元素赋值
	// 如果 i 的值超过数组下标, 会报错: 数组下标越界
	arr[i] = 10
}