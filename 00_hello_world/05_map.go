package main

import (
	"fmt"
)
/*
map 映射:
map 是一种键值对的集合, 它可以快速地根据键来查找对应的值
map 的键和值可以是任何类型, 但是键必须是不可变的
map 的键和值可以是 nil
map 的键和值可以是任何类型, 包括结构体、指针等

简单理解为 Java的 HashMap
*/
func main() {
	// 第一中声明
	var m1 map[string]string
	// 在使用 map 前, 必须先 make, make的作用就是给 map分配数据空间
	m1 = make(map[string]string, 10)
	m1["one"] = "php"
	m1["tow"] = "golang"
	m1["three"] = "java"
	fmt.Println(m1)

	// 第二种声明
	m2 := make(map[string]string)
	m2["one"] = "php"
	m2["tow"] = "golang"
	m2["three"] = "java"
	fmt.Println(m2)

	// 第三种声明
	m3 := map[string]string {
		"one": "php",
		"tow": "golang",
		"three": "java",
	}
	fmt.Println(m3)

	// map 嵌套 
	m4 := make(map[string]map[string]string)
	m4["php"] = make(map[string]string, 2)  // make(map[K]V, 容量)
	m4["php"]["id"] = "1"
	m4["php"]["desc"] = "php是世界上最好的语言"

	m4["golang"] = make(map[string]string, 2)
	m4["golang"]["id"] = "2"
	m4["golang"]["desc"] = "golang抗并发非常优秀"
	fmt.Println(m4)

	// 增删改查
	// 查找是否有 php这个子元素
	val, ok := m4["php"]
	if ok {
		fmt.Println("php 子元素存在")
		fmt.Println(val)
	} else {
		fmt.Println("php 子元素不存在")
	}
	// 修改 php 子元素的 id
	m4["php"]["id"] = "100"
	m4["php"]["desc"] = "哈哈哈"
	fmt.Println(m4)

	// 删除 php 子元素
	delete(m4, "php")
	fmt.Println(m4)

}