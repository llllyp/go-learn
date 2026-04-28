package main

import (
	"fmt"
	"math"
	"time"
)

// 定义一个结构体
type T struct {
	name string
}
func (t T) method1() {
	t.name = "new name1"
}

func (t *T) method2() {
	t.name = "new name2"
}
/*
当调用t.method1()时相当于method1(t)，实参和行参都是类型 T，可以接受。此时在method1()中的t只是参数t的值拷贝，所以method1()的修改影响不到main中的t变量。

当调用t.method2()=>method2(t)，这是将 T 类型传给了 *T 类型，go可能会取 t 的地址传进去：method2(&t)。所以 method1() 的修改可以影响 t。
*/

func main () {
	t := T{"old name"} 

	fmt.Println("method1 调用前 ",t.name)
	t.method1()
	fmt.Println("method1 调用后", t.name)

	fmt.Println("method2 调用前 ",t.name)
	t.method2()
	fmt.Println("method2 调用后", t.name)

	// 方法值 
	/*
	我们经常选择一个方法, 并且在同一个表达式里执行, 比如常见的 p.Distance()形式, 实际上就是将其分为两步来执行也是可能的
	p.Distance()叫做"选择器", 选择器会返回一个方法"值"`一个将方法(Point.Distance)`绑定到特定接收器变量的函数. 这个函数
	可以不通过指定其接收器即可被调用; 即调用时不需要指定接收器, 只要传入函数的参数即可
	*/
	r := new(Rockt)
	time.AfterFunc(10 * time.Second, func() { r.Launch() })

	// 直接用方法"值"传入AfterFunc可以更简短
	time.AfterFunc(10 * time.Second, r.Launch) // 省略了匿名函数


	// 方法表达式
	/*
	和方法"值"相关的还有方法表达式。当调用一个方法时，与调用一个普通的函数相比，必须要用选择器(p.Distance)语法来指定方法的接收器。
	当T是一个类型时，方法表达式可能会写作T.f或者(*T).f，会返回一个函数"值"，这种函数会将其第一个参数用作接收器，所以可以用通常(译注：不写选择器)的方式来对其进行调用：
	*/
	p := Point{1, 2}
	q := Point{3, 4}

	distance1 := Point.Distance // 方法表达式, 是一个函数值(相当于C语言的函数指针)
	fmt.Println(distance1(p, q))
	fmt.Printf("%T\n", distance1) // %T 表示打出数据类型, 这个必须放在Printf()中使用

	distance2 := (*Point).Distance // 方法表达式, 必须传递指针类型
	fmt.Println(distance2(&p, q))
	fmt.Printf("%T\n", distance2) // %T 表示打出数据类型, 这个必须放在Printf()中使用
	
}
type Rockt struct {/* ... */}
func (r *Rockt) Launch() { /* ... */ }

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}
