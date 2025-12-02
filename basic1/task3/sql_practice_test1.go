/*
SQL语句练习,
题目1：基本CRUD操作,
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。,
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。,
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。,
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。,

题目2：事务语句,
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

*/

/*
环境需要安装gcc：https://jmeubank.github.io/tdm-gcc/articles/2021-05/10.3.0-release
接着在powershell中设置 $env:CGO_ENABLED=1; 环境变量
数据库软件可以使用 heidisql
最后再用 gorm.io/driver/sqlite 就不会有报错了
实在不行就使用纯go版本的 modernc.org/sqlite
*/

package main

import (
	"fmt"
	"log"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint
	Name  string
	Age   int
	Grade string
}

func main() {
	fmt.Println("sql_practice_test1.")

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(sqlite.Open("sql_practice_test1.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't open the sqlite db.", err)
	}

	// 创建表 students 的表
	db.AutoMigrate(&Student{})
	// // 插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	// zhangsan := Student{Name: "张三", Age: 20, Grade: "三年级"}
	// result := db.Create(&zhangsan)
	// fmt.Println("Create result:", result)
	// // 插入一条新记录，学生姓名为 "李四"，年龄为 20，年级为 "三年级"。
	// lisi := Student{Name: "李四", Age: 12, Grade: "二年级"}
	// db.Create(&lisi)

	// 查询 students 表中所有年龄大于 18 岁的学生信息。
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	// db.Debug().Where("age > ?", 18).Find(&students)
	fmt.Println("查询结果:", students)

	// 将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
	db.Save(&Student{ID: 1, Name: "张三", Grade: "四年级", Age: 10})
	db.Where("name == ?", "张三").Find((&students))
	fmt.Println("查询结果:", students)

	db.Where("age < ?", 15).Find(&students)
	fmt.Println("查询结果:", students)
	// 删除 students 表中年龄小于 15 岁的学生记录
	db.Where("age < ?", 15).Delete(&students)
	db.Find(&students)
	fmt.Println("查询结果:", students)
}
