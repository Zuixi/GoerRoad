package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// 不关心并发操作的结果，或者有其他方式收集结果，那么WaitGroup是等待一组并发操作完成的好方法
// 如果这两个条件都不成立，我建议你改用channel和select语句
func SimpleTest() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2st goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")
}

// 通常情况下，尽可能与要跟踪的goroutine就近且成对的调用Add，但有时候会一次性调用Add来跟踪一组goroutine
func SimpleTestPro() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Println("Hello from ", id)
	}

	const num = 4
	var wg sync.WaitGroup
	wg.Add(num)
	
	for i := 0; i < num; i++ {
		go hello(&wg, i + 1)
	}
	wg.Wait()
}
