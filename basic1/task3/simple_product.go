package main

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 不使用 gorm.Model 的产品结构
type SimpleProduct struct {
	ID    uint   `gorm:"primaryKey"`
	Code  string
	Price uint
}

func main() {
	// 使用不同的数据库文件
	db, err := gorm.Open(sqlite.Open("simple_test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// 迁移模式
	err = db.AutoMigrate(&SimpleProduct{})
	if err != nil {
		panic("failed to migrate schema")
	}

	// 创建
	product := SimpleProduct{Code: "SIMPLE", Price: 50}
	err = db.WithContext(ctx).Create(&product).Error
	if err != nil {
		panic("failed to create product")
	}

	// 查询
	var products []SimpleProduct
	err = db.WithContext(ctx).Find(&products).Error
	if err != nil {
		panic("failed to find products")
	}

	fmt.Printf("SimpleProduct 表中的数据：\n")
	for _, p := range products {
		fmt.Printf("ID: %d, Code: %s, Price: %d\n", p.ID, p.Code, p.Price)
	}
}