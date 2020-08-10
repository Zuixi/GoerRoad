package leetcode

func NumberOfBoomerangs(points [][]int) int {
    // 暴力枚举
    // 使用map去重
    distance := map[int]int {}
    res := 0

    for i := 0; i < len(points); i++ {
        for j := 0; j < len(points); j++ {
            // 避免两个重复的点
            if i != j {
                dis := getdistance(points[i], points[j])
                if distance[dis] > 0 {
                    // 如果这个距离出现过，那么++
                    res += distance[dis] * 2
                }
                distance[dis]++
            }
        }
        // 对于新的起点，需要重新计算
        distance = map[int]int {}
    }

    return res
}

// 计算距离
func getdistance(x, y []int) int {
    return (x[0] - y[0]) *  (x[0] - y[0]) +  (x[1] - y[1]) *  (x[1] - y[1]) 
}