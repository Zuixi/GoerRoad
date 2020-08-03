package basic

import (
	"fmt"
)

func TestMap() {
	// 声明
	var mymap map[int]int
	mymap = make(map[int]int, 0)
	for i :=0; i < 4; i++ {
		mymap[i + 1]++
	}
	fmt.Println("include ", len(mymap))
}