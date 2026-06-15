package main

import (
	"database/sql"
	"fmt"
	"log"
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

	// 1. 更新单条记录（根据主键）
	var user User
	mysqlDb.First(&user, 1) // 先查询出记录
	user.Name = "Jinzhu Updated"
	user.Age = 22
	mysqlDb.Save(&user) // 保存更新，会更新所有字段

	// 2. 更新指定字段（只更新指定的字段）
	mysqlDb.Model(&User{}).Where("id = ?", 1).Update("name", "Jinzhu Modified")

	// 3. 更新多个字段（使用 map）
	mysqlDb.Model(&User{}).Where("id = ?", 1).Updates(map[string]interface{}{
		"name": "Jinzhu Map Update",
		"age":  25,
	})

	// 4. 更新多个字段（使用结构体）
	mysqlDb.Model(&User{}).Where("id = ?", 1).Updates(User{
		Name: "Jinzhu Struct Update",
		Age:  26,
	})

	// 5. 更新所有匹配的记录
	result := mysqlDb.Model(&User{}).Where("age > ?", 18).Update("activated_at", time.Now())
	fmt.Printf("更新了 %d 条记录\n", result.RowsAffected)

	// 6. 更新选中的字段（使用 Select 指定要更新的字段）
	mysqlDb.Model(&User{}).Where("id = ?", 1).Select("name").Updates(map[string]interface{}{
		"name":  "Selected Update",
		"age":   30, // 这个字段不会被更新，因为不在 Select 中
		"email": "updated@example.com",
	})

	// 7. 更新除指定字段外的其他字段（使用 Omit）
	mysqlDb.Model(&User{}).Where("id = ?", 1).Omit("email", "created_at").Updates(map[string]interface{}{
		"name":       "Omit Update",
		"age":        31,
		"email":      "will_not_update@example.com", // 不会被更新
		"created_at": time.Now(),                    // 不会被更新
	})

	// 8. 使用 UpdateColumn 更新单个列（不触发钩子）
	mysqlDb.Model(&User{}).Where("id = ?", 1).UpdateColumn("age", gorm.Expr("age + ?", 1))

	// 9. 使用 Raw SQL 更新
	mysqlDb.Exec("update user set age = age + 1 where id = ?", 1)

	// 10. 更新关联数据（假设有关联模型）
	// mysqlDb.Model(&user).Association("Orders").Append(&Order{Product: "New Product"})

	log.Println("更新操作完成")
}

func ConnectMysql(host, port, user, pass, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
