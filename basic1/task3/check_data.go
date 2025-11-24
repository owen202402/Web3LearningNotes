package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var products []Product
	result := db.Find(&products)
	if result.Error != nil {
		panic("failed to query products")
	}

	fmt.Printf("Found %d products:\n", len(products))
	for _, product := range products {
		fmt.Printf("ID: %d, Code: %s, Price: %d, CreatedAt: %v\n", 
			product.ID, product.Code, product.Price, product.CreatedAt)
	}
}