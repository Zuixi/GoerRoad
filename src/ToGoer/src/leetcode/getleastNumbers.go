package leetcode

import (
	"sort"
)
// #1594
func getLeastNumbers(arr []int, k int) []int {
	// 这种执行效率太低了
	// tn = o(nlogn)
	// sn = o(logn)
    sort.Ints(arr)
    return arr[:k]
}

func getLeastNumbers2(arr []int, k int) []int {
    return []int{}
}