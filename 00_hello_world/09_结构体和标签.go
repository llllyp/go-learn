package main

import "fmt"

/*
结构体 = 把多个不同类型的变量打包在一起, 自定义符合类型
- Java: 用 class 定义对象
- Go: 用 struct 定义对象
// go
type User struct {
	Name string
	Age int
}
// java
class User {
	String name;
	int age;
}

标签 Tag -> java 注解
Java：@JsonProperty("name") 注解
Go：`key:"value"` 标签

需要通过反射解析
*/

type User struct {
    Name string `json:"name" db:"user_name" validate:"required"`  // Java 属性注解 @JsonProperty("name")
    Age  int    `json:"age" db:"user_age"`
}

// 自定义tag -> Java 类注解
type User2 struct {
    // 👇 这一行 = 类上注解！固定字段，约定俗成
    TableName struct{} `db:"table:t_user"`  // Java 类注解 @Table(name = "t_user")

    // 普通字段
    ID   int
    Name string
}


func main() {
	// 1. 键值对初始化(最常用)
	user1 := User{Name: "张三", Age: 18}
	// 2. 顺序初始化
	// user2 := User{"李四", 19}
	// 3. 指针初始化(推荐)
	// user3 := &User{Name: "王五", Age: 20}

	// 访问字段
	fmt.Println(user1.Name) // 获取
	user1.Age = 21  // 修改
}
