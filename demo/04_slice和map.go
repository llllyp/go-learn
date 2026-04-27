package main

/*
slice 切片:
切片是一种引用类型, 它指向一个数组, 但是切片本身不占用内存空间
切片的长度和容量是可变的, 可以在运行时动态调整
切片的元素是数组的引用, 所以切片的元素是可变的
切片的元素可以是任何类型, 包括结构体、指针等
切片的元素可以是 nil, 但是切片本身不能是 nil

简单理解为 Java的 ArrayList

定义: var identifier []type

	var sliceA []int

---

map 映射:
map 是一种键值对的集合, 它可以快速地根据键来查找对应的值
map 的键和值可以是任何类型, 但是键必须是不可变的
map 的键和值可以是 nil
map 的键和值可以是任何类型, 包括结构体、指针等

简单理解为 Java的 HashMap
*/
func main() {

	// 定义切片
	var slice1 []int = make([]int, 10)
	// 简写
	slice2 := make([]string, 5)

}
