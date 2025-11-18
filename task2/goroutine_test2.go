/*
✅Goroutine
1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。

2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

package main

import (
	"fmt"
	"time"
)

func print1() {
	fmt.Println("task2 - goroutine_test2 - print1")
	for i := 1; i <= 10; i += 2 {
		fmt.Println("i:", i)
	}
	fmt.Println("task2 - goroutine_test2 - print1 end")
}

func print2() {
	fmt.Println("task2 - goroutine_test2 - print2")
	for i := 2; i <= 10; i += 2 {
		fmt.Println("i:", i)
	}
	fmt.Println("task2 - goroutine_test2 - print2 end")
}

func print3() {
	fmt.Println("task2 - goroutine_test2 - print3")
	for i := 11; i <= 20; i += 2 {
		fmt.Println("i:", i)
	}
	fmt.Println("task2 - goroutine_test2 - print3 end")
}

type TestFuncType func()

func main() {
	fmt.Println("task2 - goroutine_test2")

	var testFuncs []TestFuncType = []TestFuncType{print1, print2, print3}

	for _, testFunc := range testFuncs {
		go func() {
			start := time.Now()
			testFunc()
			elapsed := time.Since(start)
			fmt.Println("task using", elapsed)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("task2 - goroutine_test2 end")
}
