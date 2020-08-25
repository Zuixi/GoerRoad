package main

import (
	"fmt"
	"ToGoer/src/basic" 
	"ToGoer/src/leetcode"
	"ToGoer/src/concurrency"
	"ToGoer/src/codinginterview"
)
// package path  
// work vm is ToGoer/src/basic
// local vm is GoerRoad/src/ToGoer/src/leetcode
//  "program": "${workspaceFolder}\\GoerRoad\\src\\ToGoer\\src\\main" launch.json
// ${workspaceRoot}/src
// ${workspaceFolder}/ToGoer/src/main
func main() {
	concurrency.ActivateMultiRoutine()
	concurrency.RangeChannel()
	concurrency.CloseChannel()
	concurrency.TestResultOfChan()
	concurrency.TestSendData()
	concurrency.TestPoolWithCompare()
	concurrency.TestPool()
	concurrency.TestOnceWithTwoFunc()
	concurrency.TestOnceWithChan()
	concurrency.TestCond()
	codinginterview.NumWays(3)
	fmt.Println(codinginterview.RecurbFib(95))
	codinginterview.Fib(95)
	basic.TestSlice()
	basic.TestStruct()
	concurrency.TestRWMutex()
	concurrency.TestMutex()
	basic.TestEnum()
	basic.TestConstant()
	concurrency.SimpleTestPro()
	concurrency.SimpleTest()
	concurrency.TestJoinScopeProModify()
	concurrency.TestJoinScopePro()
	concurrency.TestJoinScope()
	concurrency.CreateJoinKnot()
	words := []string {"kuvp","q"}
	order := "ngxlkthsjuoqcpavbfdermiywz"
	leetcode.IsAlienSorted(words, order)
	points := [][]int {{0,0}, {1,0}, {2,0}}
	leetcode.NumberOfBoomerangs(points)
	leetcode.MaxNumberOfBalloons("loonbalxballpoon")
	dup := []int {1,0,2,3,0,4,5,0}
	leetcode.DuplicateZeros(dup)
	leetcode.GetNoZeroIntegers(1010)
	stones := []int {2, 7 , 4 , 1 , 8 , 1}
	leetcode.LastStoneWeight(stones)
	leetcode.RotatedDigits(678)
	leetcode.IsAnagram("anagram", "nagaarm")
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