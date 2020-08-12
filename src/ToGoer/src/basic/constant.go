package basic

import "fmt"

// 常量不是变量
// 常量只存在编译时期

func TestConstant() {

	// -------------
	// 声明和初始化
	// -------------

	// 常量是可以给定类型或者不给定类型的
	// 如果没有给定类型，那么编译器将会隐式转换

	// 没有给定类型的常量
	const ui = 12345   // kind: integer
	const uf = 3.12232 // kind: floating-point

	fmt.Println(ui)  
	fmt.Println(uf)

	// 在声明的时候给定常量类型，那么这个精度会受到限制
	const ti int = 12345    // type：int
	const tf float64 = 3.141592    // type:
	fmt.Println(ti)
	fmt.Println(tf)

	// cosnt变量需要注意值溢出的可能性，这种情况下是无效的
	// cosnt uint8 = 1000 这里将会值溢出

	// 常量运算支持很多种类
	// 常量运算的类型都是隐式转换的

	var res = 3 * 0.232 // KindFloat(3) * KindFloat(0.232)
	fmt.Println(res)

	// 常量third 是floating point kind
	const third = 1 / 3.0 //KindFloat(1) / KindFloat(3.0)
	fmt.Println(third)

	// 常量zero是 整形
	const zero = 1 / 3 // KindInt(1) / KindInt(3)
	fmt.Println("zero = ", zero)

	// 如果常量运算在一个已经声明类型和一个未声明类型之间
	// 那么未声明类型的常量将会以已经声明类型的常量作为隐式转换
	const one int8 = 1
	const two = one * 2 // int8(2) * int8(1)
	fmt.Println("one is ", one, "two is ", two)

	// 在64位机器上大的整形常量为
	const maxInt =  9223372036854775807
	fmt.Println("maxInt is ", maxInt)

	// ------
	// iota
	// ------

	const (
		myOne = iota // 0: 以0作为开始
		myTwo = iota // 1: 自增
		myThree = iota // 2:自增
	)

	fmt.Println("1. ", myOne, myTwo, myThree)

	// 在iota后面加上一个数字，是的以数字作为开始递增
	const (
		myOne1 = iota + 1 // 0: 以0 + 1作为开始
		myTwo2 // 2: 自增
		myThree3 // 3:自增
	)

	fmt.Println("2. ", myOne1, myTwo2, myThree3)

	const (
		Ldate         = 1 << iota //  1 : Shift 1 to the left 0.  0000 0001
		Ltime                     //  2 : Shift 1 to the left 1.  0000 0010
		Lmicroseconds             //  4 : Shift 1 to the left 2.  0000 0100
		Llongfile                 //  8 : Shift 1 to the left 3.  0000 1000
		Lshortfile                // 16 : Shift 1 to the left 4.  0001 0000
		LUTC                      // 32 : Shift 1 to the left 5.  0010 0000
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}