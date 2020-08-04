package main

import (
	"fmt"
	"ToGoer/src/basic" 
    "ToGoer/src/leetcode"
)
// package path  
// work vm is ToGoer/src/basic
// local vm is GoerRoad/src/ToGoer/src/leetcode
func main() {
	 res :=  []string {"leetcode","et","code"}
	leetcode.StringMatching(res)
	candies := []int {1, 1, 2, 3}
	leetcode.DistributeCandies(candies)
	basic.TestMap()
	basic.TestMultiArray()
	basic.TestArray()
	leetcode.ReverseOnlyLetters("ab-cd")
	leetcode.DetectCapitalUse("jjj");
	basic.TestAndNot()
	fmt.Println("leetcode")
	var array = []int {1,0,0,0,1,1}
	leetcode.FindMaxConsecutiveOnes(array)
	fmt.Println(leetcode.BinaryGap(22))
	fmt.Println(leetcode.GetStringOfInt(22, 2))
	fmt.Println("hello")
	basic.TestInt()

	fmt.Println()
	basic.TranferIntToString()
	basic.OperateString()

	//create txt file and writesomething in it
	var filePath = "D:/Documents/GoFiles/hello.txt"
	isSuccess := basic.ReadStreamFromFile(filePath)
	fmt.Println("is success? --> ", isSuccess)
}