package codinginterview

func TwoSum(nums []int, target int) []int {
    // n * log(n)
    // 递增排序 --> 二分?
    res := make([]int,2)
    flag := false

    for i,_ := range(nums) {
        temp := target - nums[i]

        left := i + 1
        right := len(nums) - 1
        if left >= len(nums) || flag == true{
            break
        }
        
        for {          
			if left > right {
				break
			}

            mid := left + (right - left) / 2          
            if nums[mid] == temp {
                res[0] = nums[i]
                res[1] = temp
                flag = true
                break
            }
            if nums[mid] < temp {
                left = mid + 1
            }
            if nums[mid] > temp {
                right = mid - 1
            }
        }
    }

    return res
}

func SumPro(nums []int, target int) []int {
	// 使用双指针
	// 单调有序
	var res []int = []int{0,0}
    var left, right = 0, len(nums) - 1

    for {
        if left < right {
            cur := nums[left] + nums[right]
            if cur < target {
                left++
            }
            if cur > target {
                right--
            }
            if cur == target {
                res[0] = nums[left]
                res[1] = nums[right]
                return res
            }
        } else {
            break
        }
    }
    
    return res
}