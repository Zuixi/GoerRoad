package concurrency

import (
	"fmt"
	"sync"
	"time"
	"math"
	"os"
	"text/tabwriter"
)

// mutex 表示互斥  mutual exclusion
// 互斥提供了一种并发安全的方式来表示对共享资源访问的独占

// 通过两个goroutine，试图增加和减少一个公共值，并用mutex进行访问
func TestMutex() {
	var count int
	var lock sync.Mutex

	increment := func() {
			lock.Lock()					// 在访问资源的时候加锁
			defer lock.Unlock()			// 在不需要资源的时候释放锁
			count++
			fmt.Printf("incrementing: %d\n", count)
		}		

	decreament := func() {
		lock.Lock()					// 在访问资源的时候加锁
		defer lock.Unlock()			// 在不需要资源的时候释放锁
		count--
		fmt.Printf("decrementing: %d\n", count)
	}

	// increment
	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)

		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// decrement
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)

		go func() {
			defer arithmetic.Done()
			decreament()
		}()
	}

	// 等待所有goroutine结束
	arithmetic.Wait()
	fmt.Println("Arithmetic Done")
}

// 生产者消费者模式
// 多个并发的程序不是都需要读取和写入的，对于这种考虑，我们可以使用另一个类型的互斥锁sync.RWMutex
// RWMutex 可以拥有更多锁定控制方式，可以请求锁定进行读取
// 使用RWMutex，除非有其他并发程序在进行写操作，否则就可以拥有读权限
func TestRWMutex() {

	producer := func(wg *sync.WaitGroup, lock sync.Locker) {
		defer wg.Done()
		
		for i := 5; i >= 0; i-- {
			lock.Lock()
			lock.Unlock()
			time.Sleep(1)    // 2
		}
	}

	observer := func(wg *sync.WaitGroup, lock sync.Locker) {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)

		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
	
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw, "%d\t%v\t%v\n", count,
			test(count, &m, m.RLocker()), test(count, &m, &m),
		)
	}
}