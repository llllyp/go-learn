package main

import (
	// "context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
create table user
(
    id            int auto_increment
        primary key,
    name          varchar(50)  default ''                null comment '用户名',
    email         varchar(255) default ''                not null comment '邮箱',
    age           tinyint      default 0                 not null comment '年龄',
    birthday      datetime     default CURRENT_TIMESTAMP not null comment '生日',
    member_number varchar(50)                            null comment '成员编号',
    activated_at  datetime                               null comment '激活时间',
    created_at    datetime     default CURRENT_TIMESTAMP not null,
    updated_at    datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at    datetime                               null,
    constraint u_email
        unique (email)
)
    comment '用户表' charset = utf8mb3;
*/
type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ignored      string
}
func (u *User) TableName() string {
	return "user"
}

func main() {
	mysqlDb, err := ConnectMysql("localhost", "3306", "root", "123456", "gorm")
	if err != nil {
		panic(err.Error())
	}
	user := User{Name: "Jinzhu", Email: fmt.Sprintf("%d@123.com", time.Now().UnixNano()), Age: 19, Birthday: time.Now()}
	// ctx := context.Background()
	// err := gorm.G[User](mysqlDb).Create(ctx, &user)
	result := mysqlDb.Create(&user) // 通过指针创建

	log.Println("result error: ", result.Error)
	log.Println("result rows: ", result.RowsAffected)

	// 批量插入
	users := []*User{
		{Name: "zhangsan", Email: fmt.Sprintf("%d@123.com", time.Now().UnixNano()), Age: 18, Birthday: time.Now()},
		{Name: "lisi", Email: fmt.Sprintf("%d@123.com", time.Now().UnixNano()), Age: 18, Birthday: time.Now()},
	}
	mysqlDb.Create(&users)

	// 根据 map创建
	// 通过 map[string]interface{} 与 []map[string]interface{}{}来创建记录
	mysqlDb.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "WangWu", "Email": fmt.Sprintf("%d@123.com", time.Now().UnixNano()), "Age": 17},
		{"Name": "ZhaoLiu", "Email": fmt.Sprintf("%d@123.com", time.Now().UnixNano()), "Age": 17},
	})

	
}

func ConnectMysql(host, port, user, pass, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, pass, host, port, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}