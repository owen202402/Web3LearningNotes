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
接着在powershell中设置 $env:CGO_ENABLED=1; 环境变量
数据库软件可以使用 heidisql
最后再用 gorm.io/driver/sqlite 就不会有报错了
实在不行就使用纯go版本的 modernc.org/sqlite
*/

/*

问题： SQL扩展库 是指单独的 库还是只  gorm的raw的使用？

*/

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Employees struct {
	ID        uint
	Name      string
	Deparment string
	Salary    uint
}

func main() {
	fmt.Println("introduction_to_sqlx_test1.go")

	db, err := gorm.Open(sqlite.Open("introduction_to_sqlx_test1.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("can not open sqlite db.")
	}

	// 创建表和数据
	db.AutoMigrate(&Employees{})
	employees := []*Employees{
		{Name: "张三", Deparment: "技术部", Salary: 1000},
		{Name: "李四", Deparment: "技术部", Salary: 20100},
		{Name: "王五", Deparment: "售后部", Salary: 1500},
		{Name: "赵六", Deparment: "销售部", Salary: 900},
	}
	db.Create(employees)

	// 使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	var employees1 []Employees
	db.Raw("SELECT * FROM employees WHERE deparment = ?", "技术部").Scan(&employees1)
	fmt.Println(employees1)

	// 使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var employees2 []Employees
	db.Raw("SELECT * FROM employees ORDER BY salary DESC LIMIT 1").Scan(&employees2)
	fmt.Println(employees2)
}
