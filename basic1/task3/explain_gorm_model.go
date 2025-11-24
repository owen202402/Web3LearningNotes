package main

import (
	"fmt"
	"gorm.io/gorm"
)

// gorm.Model 的定义（简化版）
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt gorm.DeletedAt `gorm:"autoCreateTime"`
	UpdatedAt gorm.DeletedAt `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 当你使用 gorm.Model 时，相当于：
type ProductWithGormModel struct {
	gorm.Model  // 这会自动添加以下字段：
	// ID        uint           // 主键
	// CreatedAt time.Time      // 创建时间
	// UpdatedAt time.Time      // 更新时间
	// DeletedAt gorm.DeletedAt // 软删除时间
	Code  string
	Price uint
}

// 如果不使用 gorm.Model，可以自定义：
type ProductWithoutGormModel struct {
	ID    uint   `gorm:"primaryKey"`
	Code  string
	Price uint
}

func main() {
	fmt.Println("gorm.Model 包含以下字段：")
	fmt.Println("1. ID (uint) - 主键，自动递增")
	fmt.Println("2. CreatedAt (time.Time) - 记录创建时间，自动设置")
	fmt.Println("3. UpdatedAt (time.Time) - 记录更新时间，自动更新")
	fmt.Println("4. DeletedAt (gorm.DeletedAt) - 软删除时间，支持软删除功能")
	
	fmt.Println("\n这就是为什么你的数据库中有 CreatedAt、UpdatedAt、Deleted_at 等字段的原因！")
}