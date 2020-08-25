package concurrency

import (
	"fmt"
	"sync"
)

// ------------
// channel 通道
// ------------

// channel可用于同步内存的访问，也可以用于goroutine之间的通信
// channel在使用的时候，你可以将一个值传递给chan，然后在其他地方取出来
// channel的声明-->
// var dataChannel  chan interface{}
// dataChannel := make(chan interface{})
// 在dataChannel中可以写入或者读取任意类型的值，这里使用了空接口
// channel也可以声明为单方向通道，仅仅只吃支持发送或者接受信息

// 声明一个只读channel, 简单使用<-符号
// var dataChannel chan intrerface{} 
// ddataChannel = make(<-chan interface{})


// 声明一个只能被发送的channel， <-放在chan关键字的右侧就行
// var dataChannel chan<- interface{}
// dataChannel = make(chan<- interface{})

// 不会经常看到实例化的单向通道，但是单向通道被用作函数参数或者返回类型是很有用的
// 或者将双向通道隐式转换为单向通道
// var receiveChan <-chan interface{}
// var sendChan chan<- interface{}
// dataChannel := make(chan interface{})
// receiveChan = dataChannel
// sendChan = dataChannel

// 需要注意的Channel是有类型的，例如string和int等类型 

func TestSendData() {
	// channel 只能传送string
	stringChannel := make(chan string)

	go func() {
		// 将字符串放入Channel
		stringChannel <- "Hello Channel"
	}()

	// 从stringChannel中读取数据并且输出
	fmt.Println("string of channel is ", <-stringChannel)
}

// -----------------
// channel的阻塞机制
// -----------------
// 仅简单的定义一个goroutine不能保证在main goroutine之前退出
// channel是包含阻塞机制的，意味着试图写入已满的channel的任何goroutine都会等待channel被清空
// 并且任何尝试从空chan读取的goroutine都将等待，直到有一个元素被放置
// goroutine尝试在channel中写入数据，所以写入成功之前不会退出
// 这两种case容易造成死锁

//-------------
// channel 死锁
// ------------
// stringChannel := make(chan string)
// go func() {
//	   if true {
//	       return
//     }
//     这里永远不会得到执行，一直等待
//	   stringChannel <- "Hello"
// }()
// 这里将永远等待，没有结果
// fmt.Println(<-stringChannel)



// ---------------------
// <-的返回值可以有两个
// ---------------------

func TestResultOfChan() {
	stringChan := make(chan string)

	go func(){
		stringChan<- "Hello, channel"
	}()

	// 这里可以看到stringChan的多个返回值
	res, ok := <-stringChan
	fmt.Printf("(%v): %v\n", ok, res)
}

// ------------------
// close channel
// ------------------

func CloseChannel() {
	stringChan := make(chan string)
	close(stringChan)

	res, ok := <-stringChan
	// 虽然关闭了channel，但是还可以读取值
	fmt.Printf("(%v): %v\n", ok, res)
}

// --------------------
// range channel
// --------------------

func RangeChannel() {
	intChan := make(chan int)

	go func() {
		// 在通道退出之前保证正常关闭
		defer close(intChan)

		for i := 1; i <= 5; i++ {
			intChan<- i
		}
	}()

	for integer := range intChan {
		fmt.Printf("%v\t", integer)
	}
}

// --------------------
// 一次解除多个goroutine
// --------------------

func ActivateMultiRoutine() {
	begin := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 这里对begin进行读取，但是由于begin中没有任何值，所以会一直阻塞
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Printf("Unblocking goroutines...\n")
	// 关闭begin 解除多个goroutine的阻塞
	close(begin)
	wg.Wait()
}