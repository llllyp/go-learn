package main

import (
	"fmt"
	"math"
)

/*
在Go语言中，接口（interface）用来描述一组方法签名，任何类型只要实现了接口中声明的全部方法，就被视为满足该接口
Go不要求显式声明“implements”，这被称为隐式实现
*/

// 定义一个形状接口
type Shape interface {
	Area() float64 // 面积
	Perimeter() float64 // 周长
}

// 实现接口(隐式实现) 长方形 实现 形状(Shape) 接口
type Rectangle struct {
	Width, Height float64
}

// 实现 Shape 的全部方法, 即视为实现 Shape
func(r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func(r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义一个圆形, 实现 形状(Shape)接口
type Circle struct {
	Radius float64 // 半径
}
// 实现 Shape 的全部方法, 即视为实现 Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 接口作为参数
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rectangle := Rectangle{10.2, 15.3}
	PrintShapeInfo(rectangle)

	circle := Circle{10}
	PrintShapeInfo(circle)
}
