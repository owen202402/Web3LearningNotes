/*
✅Channel
1. 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
   - 考察点 ：通道的基本使用、协程间通信。
2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
   - 考察点 ：通道的缓冲机制。
*/

// 疑问，既然channel有阻塞功能同步机制，这个地方要注意什么，一般会在哪里有问题

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("chanel test2.")
	// ch := make(chan int)
	ch := make(chan int, 5)
	// ch := make(chan int, 115)

	go func(ch chan<- int) {
		for i := 1; i <= 100; i += 2 {
			fmt.Printf("生产者准备发送: %d [时间: %v]\n", i, time.Now().Format("15:04:05"))
			ch <- i
			ch <- (i + 1)
			fmt.Printf("生产者发送完成: %d [时间: %v]\n", i, time.Now().Format("15:04:05"))
		}
		close(ch)
	}(ch)
	go func(ch <-chan int) {
		for c := range ch {
			fmt.Printf("消费者收到: %d [时间: %v]\n", c, time.Now().Format("15:04:05"))
			fmt.Println("recv: ", c)
			time.Sleep(1 * time.Second)
			fmt.Printf("消费者处理完成: %d [时间: %v]\n", c, time.Now().Format("15:04:05"))
		}
	}(ch)

	time.Sleep(100 * time.Second)
}
