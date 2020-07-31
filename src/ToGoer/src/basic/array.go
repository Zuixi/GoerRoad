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
		fmt.Print("%d--%d\t", index, val)
	}
	fmt.Println()
}

