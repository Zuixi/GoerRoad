package basic

import (
	"fmt"
)

// 数组长度固定
func TestArray() {
	// 声明数组
	var arrOne [5]int

	// 遍历数组，并且输出索引和元素
	for index, val := range arrOne {
		fmt.Printf("%d--%d\t", index, val)
	}
	fmt.Println()

	// 仅仅打印元素
	for _, val := range arrOne {
		fmt.Printf("%d\t", val)
	}
	fmt.Println()

	// 默认情况下数组的每个元素都会被初始化对应的0值
	var zero = [2]int{1}
	fmt.Println("default int value is ", zero[1])

	// 如果在声明数组长度的位置使用...代替，那么这个数组的长度是根据初始化的值的个数来计算的
	var three = [...]int {1, 2, 3}
	fmt.Println("length of three is ", len(three))

	// 比较数组相等
	// 长度一样，元素类型一样才可以进行比较
	a := [2]int{1,2}
	b := [2]int{1,2}
	c := [2]int{1,3}
	fmt.Println("a == b is ", a == b, "\ta == c is ", a == c, "\tb == c is ", b == c)
}

// go语言的多维数组
func TestMultiArray() {
	// 声明语法示例 --> var arrayName [size1][size2][...][sizen] arrat_type

	// 以二维数组作为例子
	var array [4][2]int    // 维度分别是4和2
	array = [4][2]int {{1,2}, {2,3},{3,4},{4,5}}
	fmt.Println("在二维数组中, len(array) is ", len(array))

	// 声明并初始化数组中索引为1和3的元素
	var myarr = [4][2]int {1: {1,1}, 3:{3,3}}
	fmt.Println("myarr should be ", myarr[1][1])
	fmt.Println()
	//
}

