/*
✅锁机制
1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
   - 考察点 ： sync.Mutex 的使用、并发数据安全。
2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
   - 考察点 ：原子操作、并发数据安全。
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("lock test1.")
	var mu sync.Mutex
	var count int = 0

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("count: ", count)
}
