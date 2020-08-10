package leetcode

// #1184
func DistanceBetweenBusStops(distance []int, start int, destination int) int {

    // 只有两种走法，一种是逆时针，一种是顺时针，直接比较最小值返回即可
    // 1. 顺时针
    if destination == start {
        return 0
    }
    
    if destination < start  {
        destination,start = start, destination
    }

    sum := 0
    for i := start; i < destination; i++ {
        sum += distance[i]
    }

    res := 0
    for _, val := range(distance) {
        res += val
    }
    if res - sum > sum {
        return sum
    } else {
        return res - sum
    }
}