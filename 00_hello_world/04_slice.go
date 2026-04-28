package main

import (
	"fmt"
)



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

*/
func main() {

	// testSlice2()
	// emptySlice()
	// testSubSlice()
	appendAndCopy()
}

/*
定义切片
*/
func testSlice1() {
	// 定义数组
	var arr [10]int
	// 定义切片
	var slice1 []int = make([]int, 10) // len 10 是切片的初始长度
	fmt.Println(slice1)
		
	// 简写
	slice2 := make([]string, 5)
	fmt.Println(slice2)

	// 初始化切片
	s := []int{1, 2, 3, 4} // cap=len=4
	fmt.Println(s)

	s2 := arr[:] // 初始化s2切片是数组arr的引用
	fmt.Println(s2)

	// 将 arr 下标index从 start 到 end 部分创建一个新的切片
	// s3 := arr[startIndex:endIndex]
	// 缺省 start 到最后一个元素
	// s4 := arr[startIndex:]
	// 从第一个元素开始
	// s5 := arr[:endIndex]
	// 通过切片 初始化 切片
	s6 := s2[1:5]
	fmt.Println(s6)

	// ######################## 内置函数
	// len() 返回切片的长度
	fmt.Println(len(s6))
	// cap() 返回切片的容量
	fmt.Println(cap(s6))

}
/*
切片内置函数
*/
func testSlice2() {
	var slicce = make([]int, 3, 5)
	fmt.Printf("len = %d cat = %d slice=%v\n", len(slicce), cap(slicce), slicce)
}

func emptySlice() {
	var emptySlice []int
	if emptySlice == nil {
		fmt.Println("切片是空的")
	}
	fmt.Printf("len = %d cat = %d slice=%v\n", len(emptySlice), cap(emptySlice), emptySlice)
}

/*
切片截取
*/
func testSubSlice() {
	// 创建切片
	numbers := []int{1,2,3,4,5,6,7,8,9}

	// 打印原始切片
	fmt.Println("numbers = ", numbers)

	// 打印切片从索引1(包含)到4(不包含)
	fmt.Println("numbers[1:4] = ", numbers[1:4])

	// 默认下限为0
	fmt.Println("numbers[:4] = ", numbers[:4])

	// 默认上限为切片长度
	fmt.Println("numbers[4:] = ", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	// 打印 子切片从索引 0 到索引 2
	number2 := numbers1[:2]
	printSlice(number2)

	// 打印 子切片从索引 2 到索引 5
	number3 := numbers1[2:5]
	printSlice(number3)
}

/*
append() 和 copy() 函数
append() 函数用于在切片末尾添加元素
copy() 函数用于将一个切片复制到另一个切片

如果想增加切片的容量, 必须创建一个新的更大的切片, 并把原分片的内容拷贝到新切片中
*/
func appendAndCopy() {
	var numbers []int
	printSlice(numbers)

	// 允许追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)

	// 向切片添加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers)

	// 同时添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	// 创建切片 number1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	printSlice(numbers1)

	// 拷贝 numbers 的内容到 numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cat = %d slice=%v\n", len(s), cap(s), s)
}
