package leetcode

func DistributeCandies(candies []int) int {
    var cat map[int]int
    cat = make(map[int]int, 0)
    for _, val := range candies {
        cat[int(val)]++
    }
    var s = len(cat)
    if len(candies) / 2 < s {
        return len(candies) / 2
    }
    return len(candies) / 2
}