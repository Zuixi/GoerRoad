package main

import (
	"fmt"
	"GoerRoad/src/ToGoer/src/basic" 
	"GoerRoad/src/ToGoer/src/leetcode"
)

func main() {
	fmt.Println("leetcode")
	var array = []int {1,0,1,1,0,1}
	leetcode.FindMaxConsecutiveOnes(array)
	fmt.Println(leetcode.BinaryGap(22))
	fmt.Println(leetcode.GetStringOfInt(22, 2))
	fmt.Println("hello")
	basic.TestInt()

	fmt.Println()
	basic.TranferIntToString()

	fmt.Println("Generate rand string as follows")
	res := basic.GenerateRandString()
	fmt.Println("return string is ", res)

	basic.OperateString()

	//create txt file and writesomething in it
	var filePath = "D:/Documents/GoFiles/hello.txt"
	isSuccess := basic.ReadStreamFromFile(filePath)
	fmt.Println("is success? --> ", isSuccess)
}