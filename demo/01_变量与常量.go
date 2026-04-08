package main

import (
	"fmt"
	"unsafe"
)

var x, y int
var ( // 这种分解的写法, 一般用于全局变量
	a = 10
	b = 20
)

var c, d int = 1, 2

var e, f = 3, "4"

func main() {
	// 1.定义变量 1
	// 1.1 指定变量类型
	var v1 int
	fmt.Println(v1)
	// 1.2 根据值自动判断变量类型
	var v2 int = 10
	fmt.Println(v2)
	// 1.3 省略 var, := 左边的变量不能是已经生命过的 只能在函数体内使用
	v3 := 20
	fmt.Println(v3)

	_, value := 100, 200 // 实际上 100 被废弃, 不具备读特性
	fmt.Println(value)

	fmt.Println(x, y, a, b, c, d, e, f, value)

	// 2. 定义常量 格式: const 变量名 type = 值
	const c1 string = "abc"
	// 隐式定义
	const c2 = "bbb"
	const c3, c4, c5 = 999, false, "888" // 多重赋值

	const length, width = 10, 5
	area := length * width
	fmt.Printf("矩形的面积为: %d\n", area)
	// 常量枚举
	const (
		Unknown = 0 // 未知性别
		Male    = 1 // 男性
		Female  = 2 // 女性
	)
	// 常量可以使用 len() cap() unsafe.SIzeof()
	const (
		ca = "abc"
		cb = len(ca)
		cc = unsafe.Sizeof(ca)
	)
	fmt.Println(ca, cb, cc)

	// 优雅的常量 iota
	/*
		iota 是一个常量生成器, 可以在 const 声明中使用
		iota 从 0 开始, 每个 const 声明中使用一次, 就会自动增加 1
	*/
	const(
		CategoryBooks = iota
		CategoryHealth
		CategoryClothing
	)
	fmt.Println(CategoryBooks, CategoryHealth, CategoryClothing)
	const(
		IgEggs = 1 << iota
		IgMilk
		IgNuts
		IgGluten
	)
	//
	type ByteSize float64
	const(
		_ = iota
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB // 1 << (10 * 2)
		GB // 1 << (10 * 3)	
		TB // 1 << (10 * 4)
		PB // 1 << (10 * 5)
	)
	// iota定义在公式里 在下一行增长, 而不是立即取得它的引用
	const(
		Apple, Banana = iota + 1, iota + 2
		Cherimoya, Durian
		Elerberry, Fig
	)


}
