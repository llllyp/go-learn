package main

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model // 自动添加 ID, CreatedAt, UpdatedAt 字段
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	ctx := context.Background()

	//自动建表
	db.AutoMigrate(&Product{})

	// 创建
	product := Product{Code: "D42", Price: 100}
	err = db.WithContext(ctx).Create(&product).Error
	if err != nil {
		panic(err)
	}

	// 查询
	var product1 Product
	err = db.WithContext(ctx).Where("id = ?", product.ID).First(&product1).Error
	if err != nil {
		panic(err)
	}

	var products []Product
	err = db.WithContext(ctx).Where("code = ?", "D42").Find(&products).Error
	if err != nil {
		panic(err)
	}

	// 更新 - 将产品价格更新为 200
	err = db.WithContext(ctx).Model(&product1).Update("Price", 200).Error
	if err != nil {
		panic(err)
	}

	// 更新 - 更新多个字段
	err = db.WithContext(ctx).Model(&product1).Updates(Product{Code: "D42", Price: 100}).Error
	if err != nil {
		panic(err)
	}

	// 删除 - 删除产品
	err = db.WithContext(ctx).Delete(&product1).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(products)

}
