package main

import (
	"fmt"
)



/*
Golang的语言中提供了断言的功能。golang中的所有程序都实现了interface{}的接口
所有的类型如 string, int, int64 甚至是自定义的 struct 类型都就此拥有了 interface{} 的接口，
这种做法和 java 中的 Object 类型比较类似。那么在一个数据通过func funcName(interface{})的方式传进来的时候，
也就意味着这个参数被自动的转为interface{}的类型。
*/

func funcName(a interface{}) {
	// return string(a) // 此时编译器返回 : cannot convert a (type interface{}) to type string: need type assertion

	// 此时，意味着整个转化的过程需要类型断言。类型断言有以下几种形式：
	// 1. 直接断言使用
	fmt.Println(a.(string))
	// 如果断言失败, 一般会导致 panic的发生. 为了防止 panic, 需要在断言前进行一定的判断
	value, ok := a.(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("It's not ok for type string")
	}

	// 2. 配合 switch 使用
	switch a.(type) {
	case string:
		fmt.Println("It's a string")
	case int:
		fmt.Println("It's a int")
	case bool:
		fmt.Println("It's a bool")
	default:	
		fmt.Println("It's not a string or int or bool")
	}

}

func main() {
	funcName("hello")
	funcName(100)
	funcName(true)
}