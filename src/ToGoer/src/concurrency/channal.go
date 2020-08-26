package concurrency

import (
	"fmt"
	"sync"
	"bytes"
	"os"
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
// var dataChannel chan interface{} 
// dataChannel = make(<-chan interface{})


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

// ------------------
// 缓冲通道
// ------------------

// 缓冲通道的声明
// var dataChan chan interface{}
// dataChan = make(chan interface{}, 5)
// 创建容量为4的缓冲通道, 意味着可以放置4个元素在chan上，不管是否被读取;
// 在chan数量达到上限之前，写入操作都不会阻塞
// 无缓冲channel也可以按缓冲channel定义，可以当做容量为0
// a := make(chan int)
// b := make(chan int, 0)

// 阻塞的含义
// 向一个已满的channel写入数据会出现阻塞
// 从一个空的channel读取数据会出现阻塞
// 满和空都是对于容量或者缓冲区大小而言的
// 对于无缓冲channel而言，任何写入行为之后都是满的
// 对于容量为4的缓冲channel而言，4次写入之后是满的，第五次会阻塞

// 缓冲channel在在某些情况下很有用，但是应该小心使用
// 缓冲channel很容易成为不成熟的优化，让死锁变得更隐蔽
// 下面的代码用于了解缓冲channel的工作方式

func BuffChannel() {
	var stdoutBuff bytes.Buffer
	// 确保在进程退出之前将缓冲区内容输出到标准输出
	defer stdoutBuff.WriteTo(os.Stdout)

	intChan := make(chan int, 4)
	go func() {
		defer close(intChan)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intChan <- i
		}
	}()

	for integer := range intChan {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

// --------------------
// channel的默认值 nil
// --------------------

func DefaultChan() {
	// 尝试从一个nil chan中读取值
	// 会出现deadlock fatal error: all goroutines are asleep - deadlock
	// 这是在main函数中执行才会，在goroutine中执行会阻塞
	var intChan chan int
	//fmt.Printf("default value in nil channel is %d \n", <-intChan)

	// 这个goroutine会一直阻塞，因为没有值写入
	go func() {
		fmt.Printf("default value in nil channel is %d \n", <-intChan)
	}()

	// 对一个nil channel写入数据
	// 同样会出现 deadlock fatal error: all goroutines are asleep - deadlock
	//intChan <- 4

	// 对一个nil channel关闭
	// panic: close of nil channel
	// 程序挂掉了
	close(intChan)

	// 通过对nil channel的测试，可以知道无论如何都需要确保chan在工作之前已经被初始化
}

// --------------------------------------
// 如何组织不同类型的channel来构建稳健的程序
// --------------------------------------

// 首先要确保channel在正确的环境中，也就是分配channel所有权
// 这里的所有权指的是goroutine的实例化、写入和关闭，明确哪个goroutine拥有该channel
// 一个goroutine拥有一个channel时, 应该-->
// 1. 初始化该channel
// 2. 执行写入操作或将所有权交给另外的goroutine
// 3. 关闭该channel
// 4. 将上面三件事封装在一个列表中，并通过订阅channel将其公开

// 将上面的责任分配给channel 所有者，会发生一些事情
// 初始化channel的人，需要知道写入nil channel会带来死锁的危险
// 初始化channel的人，需要了解关闭nil channel会带来panic的危险
// 决定何时关闭channel的人，需要了解写入已经关闭channel会带来panic的危险
// 决定何时关闭channel的人， 需要了解多次关闭channel会带来panic的危险
// 在编译时期，需要使用类型检查器来防止对channel进行不正确的写入

// 作为一个channel的消费者， 只需要关注几件事情-->
// 1. channel什么时候会关闭
// 2. 处理基于任何原因出现的阻塞
// 想要知道channel什么时候关闭，检查chan第二个返回值就行了；
// 对于第二点，没有统一的做法，但是作为消费者可以明确的是，读取操作可以并且必将产生阻塞

// --------------------------------------------------------
// 一个goroutine拥有一个channel和消费者, 处理阻塞问题
// --------------------------------------------------------
func DealBlockWithGoroutine() {
	chanOwner := func() <-chan int {
		resultChan := make(chan int, 5)

		go func() {
			// 确保在操作完成之后关闭channel, 这是所有者的职责
			defer close(resultChan)

			for i := 0; i <= 5; i++{
				resultChan <- i
			}
		}()
		// 这里resultChan将被隐式转换为只读channel
		return resultChan
	}

	resultChan := chanOwner()
	// 这里消费者消费channel， 只关心阻塞和关闭
	for result := range resultChan {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Printf("Done Receiving.\n")
}