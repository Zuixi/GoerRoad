package basic

import (
	"fmt"
)

// number 表示一个学生的信息，表示
type number struct {
	age  uint8
	flag int16
	id   int	
}

func TestStruct() {
	// ----------
	// 声明和初始化
	// ----------
	
	// 声明一个number的示例并且将其初始化为零值
	// 内存的分配情况如下所示-->
	// uint8 一个字节， int16 两个字节， int 4个字节
	// 总共加起来有7个字节，但实际上是8个字节
	// 因为在Go中有填充和对齐的说法
	// 由于需要对齐，那么就会在uint8和int16之间填充

	// 为何需要对齐--> 在对齐边界的内存上读取效率更高
	// 所以go语言通常会解决人们不太注意的边界对齐的问题

	// 规则一
	// 根据变量的大小，go会自动确定需要对齐的方式。
	// 每个两个字节的变量都必须遵循两个字节的边界。由于uint8只有一个字节并且在第一个位置，
	// 而int16有两个字节且出现在第二个位置，地址被跳过的字节会填充一个字节
	// 类似地，如果第二个位置有4个字节的话，那么我们会填充3个字节

	var num1 number
	// 输出变量
	fmt.Printf("%+v\n",num1)

	// 规则二:
	// 结构体的最大字段代表着整个结构体的填充
	// 需要尽可能的使填充的字节最小
	// 始终从高到低的顺序排列字段


	// 在下面的这个例子中，整个结构体的大小都必须遵守8字节
	// type xample struct {
	// 	   counter int64
	//     pi      float32
	//     flag    bool
	// }

	// 声明和初始化number的字段的数值
	// 每一行都必须以逗号结束
	num2 := number {
		age: 23,
		flag: 3,
		id:99,
	}

	// 输出num2字段的信息
	fmt.Println("age is ", num2.age)
	fmt.Println("flag is ", num2.flag)
	fmt.Println("id is ", num2.id)

	// 声明一个匿名类型的变量，并且使用结构体初始化
	student := struct {
		id int
		age uint8
		name string
	}{
		id:23,
		age:21,
		name:"www",
	}
	fmt.Println("id is ", student.id)
	fmt.Println("age is ", student.age)
	fmt.Println("name is ", student.name)

	fmt.Println("Test over\n")
}
      