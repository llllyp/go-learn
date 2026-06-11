package main

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdateAt     time.Time
	ignored      string
}

/*
- 具体数字类型如: uint, string 和 uint8 直接使用
- 指向 *string 和 *time.Time 类型的指针表示可空字段
- 来自 database/sql 包的 sql.NullString 和 sql.NullTime 用于具有更多控制的可空字段
- CreatedAt 和 UpdatedAt 是特殊字段，当记录被创建或更新时，GORM 会自动向内填充当前时间
*/
// gorm.Model 的定义
/*
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"` // 用于软删除
}
*/
/*
type User2 struct {
  Name string `gorm:"<-:create"` // 允许读和创建
  Name string `gorm:"<-:update"` // 允许读和更新
  Name string `gorm:"<-"`        // 允许读和写（创建和更新）
  Name string `gorm:"<-:false"`  // 允许读，禁止写
  Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
  Name string `gorm:"->;<-:create"` // 允许读和写
  Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
  Name string `gorm:"-"`  // 通过 struct 读写会忽略该字段
  Name string `gorm:"-:all"`        // 通过 struct 读写、迁移会忽略该字段
  Name string `gorm:"-:migration"`  // 通过 struct 迁移会忽略该字段
}
*/
