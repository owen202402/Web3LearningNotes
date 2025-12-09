/*
Sqlx入门
题目1：使用SQL扩展库进行查询
- 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
  - 要求 ：
    - 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
    - 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
题目2：实现类型安全映射
- 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
  - 要求 ：
    - 定义一个 Book 结构体，包含与 books 表对应的字段。
    - 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

/*
环境需要安装gcc：https://jmeubank.github.io/tdm-gcc/articles/2021-05/10.3.0-release
接着在powershell中设置 $env:CGO_ENABLED=1; 环境变量, 或者直接在gitbash中使用就不需要
数据库软件可以使用 heidisql
最后再用 gorm.io/driver/sqlite 就不会有报错了
实在不行就使用纯go版本的 modernc.org/sqlite
*/

/*

问题：
实现类型安全映射是什么特指吗？
-- 没有特指的话，那和题目一的差别在哪里？ 答：没有什么差别

该题目需要重写写下，使用扩展库sqlx去写

*/

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// books 表，包含字段 id 、 title 、 author 、 price 。
type Books struct {
	ID     uint
	Title  string
	Author string
	Price  uint
}

func main() {
	fmt.Println("introduction_to_sqlx_test2.go")

	db, err := gorm.Open(sqlite.Open("introduction_to_sqlx_test2.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("can not open sqlite db.")
	}

	db.AutoMigrate(&Books{})
	books := []*Books{
		{Title: "一年级课本", Author: "教育局1", Price: 21},
		{Title: "一一年级课本", Author: "教育局11", Price: 221},
		{Title: "二年级课本", Author: "教育局2", Price: 41},
		{Title: "二二年级课本", Author: "教育局22", Price: 441},
		{Title: "三年级课本", Author: "教育局3", Price: 31},
		{Title: "三三年级课本", Author: "教育局33", Price: 331},
		{Title: "四年级课本", Author: "教育局4", Price: 11},
		{Title: "四四年级课本", Author: "教育局44", Price: 111},
	}
	db.Create(books)

	var books1 []Books
	db.Raw("SELECT * FROM books WHERE price > ?", 50).Scan(&books1)
	fmt.Println(books1)

}
