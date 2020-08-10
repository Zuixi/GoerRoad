package offercode

// leetcode #1564
func FindRepeatNumber(nums []int) int {
    res := make([]int, len(nums))

    for _, val := range(nums) {
        res[val]++
        if res[val] > 1 {
            return val
        }
    }

    return 0
}