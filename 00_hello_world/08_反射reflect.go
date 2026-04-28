package main

import (
	"fmt"
	"reflect"
)
/*
# Go反射是什么
反射 = 在程序运行时，动态获取变量的【类型】和【值】，并能动态修改变量、调用方法。

Java 				反射					Go 反射	作用
Object				interface{}				所有类型的 “万能容器”
getClass()			reflect.TypeOf()		获取类型信息
getField()/invoke()	reflect.ValueOf()		获取 / 操作值信息
Field				reflect.StructField		结构体字段
Method				reflect.Method			方法

# 反射量大核心: Type 和 Value
Type: 反射中的类型信息
Value: 反射中的值信息
1. reflect.Type: 类型信息(只读)
- 变量是什么类型
- 结构体有哪些字段
- 方法有哪些
reflect.TypeOf(变量)

2.reflect.Value: 值信息(可读可写)
- 变量当前值是多少
- 动态修改值
- 动态调用方法
reflect.ValueOf(变量)

# 4个核心API
1. 获取类型, 种类
t := reflect.TypeOf(obj)
t.Name()  // 类型名（如 User、int）
t.Kind()  // 底层种类：struct/int/string/slice/ptr 等

2. 获取和修改值
v := reflect.ValueOf(&obj).Elem()  // 必须传指针才能修改！
v.FieldByName("Name").SetString("李四")

3. 读取结构体 Tag（最常用！）
field.Tag.Get("json")
field.Tag.Get("db")
field.Tag.Get("validate")

4. 动态调用方法
// 获取方法
method := v.MethodByName("SayHello")
// 构建参数
args := []reflect.Value{reflect.ValueOf("你好")}
// 调用
method.Call(args)

*/

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {

	user := User{"张三", 18}

	// 1. 获取类型 Type
	t := reflect.TypeOf(user)
	fmt.Println("类型名: ", t.Name())
	fmt.Println("种类: ", t.Kind())

	// 2.获取值 Value
	v := reflect.ValueOf(user)
	fmt.Println("值: ", v)

	// 3. 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)    //字段类型信息
		value := v.Field(i).Interface() //字段值
		fmt.Printf("字段:%s: 标签:%s, 值: %v\n", field.Name, field.Tag.Get("json"), value)
	}

}
