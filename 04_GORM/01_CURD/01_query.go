package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	Name         string         `gorm:"size:50;default:'';comment:用户名"`
	Email        string         `gorm:"size:255;default:'';not null;unique;comment:邮箱"`
	Age          uint8          `gorm:"default:0;not null;comment:年龄"`
	Birthday     time.Time      `gorm:"default:CURRENT_TIMESTAMP;not null;comment:生日"`
	MemberNumber sql.NullString `gorm:"size:50;comment:成员编号"`
	ActivatedAt  sql.NullTime   `gorm:"comment:激活时间"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP;not null;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index;comment:软删除时间"`
}

func (User) TableName() string {
	return "user"
}

func main() {
	mysqlDb, err := ConnectMysql("localhost", "3306", "root", "123456", "gorm")
	if err != nil {
		panic(err.Error())
	}

	// 1. 查询单条记录
	var user User
	// 根据主键查询
	mysqlDb.First(&user, 1)
	fmt.Printf("根据ID查询用户: %+v\n", user)

	// 根据条件查询
	mysqlDb.Where("name = ?", "Jinzhu").First(&user)
	fmt.Printf("根据名称查询用户: %+v\n", user)

	// 2. 查询多条记录
	var users []User
	// 查询所有用户
	mysqlDb.Find(&users)
	fmt.Printf("查询所有用户: %+v\n", users)

	// 查询年龄大于18的用户
	mysqlDb.Where("age > ?", 18).Find(&users)
	fmt.Printf("查询年龄大于18的用户: %+v\n", users)

	// 3. 条件查询
	// IN查询
	mysqlDb.Where("name IN ?", []string{"Jinzhu", "Mike"}).Find(&users)

	// LIKE查询
	mysqlDb.Where("name LIKE ?", "%jin%").Find(&users)

	// AND查询
	mysqlDb.Where("name = ? AND age > ?", "Jinzhu", 18).Find(&users)

	// 4. 选择指定字段
	mysqlDb.Select("name", "age").Find(&users)

	// 5. 排序
	mysqlDb.Order("age desc, name").Find(&users)

	// 6. 分页
	var pageUsers []User
	pageSize := 10
	pageNum := 1
	mysqlDb.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&pageUsers)

	// 7. 关联查询（如果有关联表）
	// mysqlDb.Preload("Orders").Find(&users)

	// 8. 软删除查询
	mysqlDb.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)

	// 9. 计数
	var count int64
	mysqlDb.Model(&User{}).Count(&count)
	fmt.Printf("用户总数: %d\n", count)
}

func ConnectMysql(host, port, user, pass, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
