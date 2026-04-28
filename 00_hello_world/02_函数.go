package main

import (
	"fmt"
)

func main() {
	// go 函数可以返回多个值
	a, b := swap("张三", "李四")
	fmt.Println(a, b)

	// init函数与import
	// init函数可在 package main中, 可在其他package中定义, 可以在同一个 package中出现多次
	// main 函数只能在 package main中

	// 执行顺序
	// 1. 全局变量初始化
	// 2. init函数执行
	// main 函数执行
	// ../doc/00基础语法/golang函数执行顺序.png

	// 函数参数
	// 函数如果使用参数, 该变量可称为函数的形参, 形参就像定义在函数体内的局部变量一样, 只能在函数体内使用
	// 实参是调用函数时传递的值, 会赋值给形参

	// 值传递
	var x, y int = 100, 200
	fmt.Println("交换前 x:", x)
	fmt.Println("交换前 y:", y)
	swap2(x, y)
	fmt.Println("swap2交换后 x:", x)
	fmt.Println("swap2交换后 y:", y)

	// 引用传递(指针传递)
	swap3(&x, &y)
	fmt.Println("swap3交换后 x:", x)
	fmt.Println("swap3交换后 y:", y)	


}

func swap(x, y string) (string, string) {
	return y, x
}

// 值传递
func swap2(x, y int) {
	var temp int
	
	temp = x // 保存 x 的值
	x = y  // 将 y 的值赋给 x
	y = temp // 将 temp 的值赋给 y
}

// 引用传递(指针传递)
func swap3(x, y *int) {
	var temp int

	temp = *x // 保存 x 地址上的值
	*x = *y // 将 y 的值赋给 x
	*y = temp // 将 temp 的值赋给 y
}
