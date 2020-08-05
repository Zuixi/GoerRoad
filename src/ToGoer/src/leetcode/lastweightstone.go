package leetcode

import "sort"

// #1046
func LastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	size := len(stones)
	// 快速排序从小到大
	sort.Ints(stones)
	for stones[size-2] != 0 {
		stones[size-1] -= stones[size-2]
		stones[size-2] = 0
		sort.Ints(stones)
	}
	return stones[size-1]
}
