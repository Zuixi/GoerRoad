package concurrency

import (
	"time"
	"fmt"
)

// channel是goroutine的粘合剂， select是channel的粘合剂
// select 可以让我们在项目中组合channel来形成更大的抽象来解决实际问题

// -------------------------
// select 的使用
// -------------------------

// 可以在单个函数或者类型定义中找到将本地channel绑定在一起的select语句
// 也可以在全局范围内连接系统级别或者多个组件的使用范例
// select还可以帮助你安全将channel和业务层面的概念(如取消、超时、等地和默认值)结合在一起

func SampleSelect() {
	var intChan <-chan interface{}
	var stringChan <-chan interface{}

	select {
	case <- intChan:
		// do something
	case <-stringChan:
		// do something
	}

	// select看起来和switch很像，都包含case分支，但是有不同点-->
	// select的分支不会被顺序测试，如果没有任何分支的条件满足，那么select会一直等待某个case语句完成；
	// 所有channel的读取和写入都会被同时考虑， 用来看他们是否准备好
	// 如果没有任何channel准备就绪，那么select会一直阻塞，直到某个channel准备好
}

func LockAndUnlockSelect() {
	start := time.Now()

	channel := make(chan interface{})
	go func() {
		// 在5s的阻塞后悔关闭channel
		defer close(channel)
		time.Sleep(10 * time.Second)
	}()

	fmt.Printf("Blocking on read...\n")
	select {
	case <-channel:
		fmt.Printf("Unblocking %v later.\n", time.Since(start))
	}
}

// select是一种简单有效的实现阻塞等待的方式，但会有以下问题-->
// 当多个channel需要读取时会发生什么
// 当所有channel都尚未初始化完成，又会发生什么
// 当我们想做点什么，但是channel没有准备好

func ReadMultiChan(){
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var countOfC1, countOfC2 uint

	// c1, c2被关闭之后，仍旧可以读取nil
	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			countOfC1++
		case <-c2:
			countOfC2++
		}
	}

	// 从这个循环中可以知道go有大概一半的时间是在c1读取，一半在c2读取
	fmt.Printf("countOfC1 is %d\t\tcountOfC2 is %d", countOfC1, countOfC2)
	// 为了是多个channel读取的机会大概相等，我们可以使用一个随机变量，已决定哪个case得到执行
}

// 如果多个channel尚未初始化，并且自己无法进行处理，并且不能持续等待下去，则可以进行超时处理
func TimeOutChan() {
	var c<-chan int
	select {
		// 这个case将会永远阻塞，因为没有值输入
	case <-c:
		// do something
	case <-time.After(5 * time.Second):
		fmt.Printf("Time out")
	}
}

// 如果自己想做点什么，但是channel没有准备好，可以使用default语句
func DoOthers() {
	start := time.Now()
	var c1, c2 <-chan int

	select {
	case <-c1:
		// do something
	case <-c2:
		// do something
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}

	// default 语句几乎是立马就得到执行，这允许我们在不阻塞的情况下退出选择快
}

// 有一个循环正在做某种工作，偶尔检查它是否应该停止
func ForSelect() {
	done := make(chan interface{})

	go func() {
		time.Sleep(4 * time.Second)
		close(done)
	}()

	workCount := 0
	
	Loop:
	for {
		select {
		case <-done:
			break Loop
		default:
			// do something
		}

		// simulate work
		workCount++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v circles of work before signalled to stop.\n", workCount)
}

// ---------------------------
// 空的select语句将会永远阻塞
// select {}
// ---------------------------