package main

import (
	"fmt"
	"math"
)

/*
当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了。可以根据选择来调用接收器各不相同的方法。
示例:
变量op代表Point类型的addition或者subtraction方法，Path.TranslateBy方法会为其Path数组中的每一个Point来调用对应的方法：
*/

type Point struct {
	X, Y float64
}

// 给 Point 定义一个方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Print() {
	fmt.Printf("{%f, %f}\n", p.X, p.Y)
}
// 定义一个 Point切片类型 Path
type Path []Point

// 方法的接收器 是Path类型的数据, 方法的选择器是TranslateBy(Point, bool)
func(path Path) TranslateBy(anather Point, add bool) {
	// 定义一个 op变量 类型是方法表达式 能够接收Add,和 Sub方法
	var op func(p, q Point) Point
	if add {
		op = Point.Add // 给 op 变量赋值为 Add
	} else {
		op = Point.Sub // 给 op 变量赋值为 Sub
	}

	for i := range path {
		path[i] = op(path[i], anather)
		path[i].Print()
	}
}


func main() {
	points := Path{
		{1, 2},
		{3, 4},
	}
	anather := Point{5, 5}

	points.TranslateBy(anather, true)

	fmt.Println("--------------------------")

	points.TranslateBy(anather, false)
}
