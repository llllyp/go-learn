package main

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mian() {
	// 想要正确的处理 time.Time ，需要带上 parseTime 参数
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbName?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//mysql驱动 一些高级配置
	mysqlDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	// PostgreSQL
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// https://github.com/go-gorm/postgres
	pgSqlDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	// clickhouse 
	ckDsn := "clickhouse://gorm:gorm@localhost:9942/gorm?dial_timeout=10s&read_timeout=20s"
	db, err := gorm.Open(clickhouse.Open(ckDsn), &gorm.Config{})

	// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
	db.AutoMigrate(&User{})
	// 设置表选项
	db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&User{})

	// 插入
	db.Create(&user)

	// 查询
	db.Find(&user, "id = ?", 10)

	// 批量插入
	var users = []User{user1, user2, user3}
	db.Create(&users)
	// ...
}
