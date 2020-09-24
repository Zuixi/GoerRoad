package concurrency

import (
	"fmt"
	"time"
)

// ----------------------
// goroutine的终止方法
// ----------------------

// 1. 当任务完成
// 2. 遇到错误无法继续任务
// 3. 被告知停止任务

// goroutine泄露的例子

func LeakGoroutine() {
	
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork Exited")
			defer close(completed)

			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	// 这里传入nil， 将会一直阻塞包含doWork的goroutine，直到程序结束
	doWork(nil)

	// 继续干其他的事情
	fmt.Println("Done.")
}

// 为了结局上面的goroutine泄露，一般解决办法是建立一个信号，父例程将该通道传递给子例程，然后在想要取消子例程时关闭该通道
func SolveLeak() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		
		go func() {
			defer fmt.Println("doWork Exited") 
			defer close(completed)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <- done:
					return
				}
			}
			
		}()
		return completed
	}

	done := make(chan interface{})
	doWork(done, nil)

	go func() {
		// 在1秒之后取消goroutine
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	fmt.Println("Done.")
}