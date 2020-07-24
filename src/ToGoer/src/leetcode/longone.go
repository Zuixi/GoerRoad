package leetcode

func FindMaxConsecutiveOnes(nums []int) int {
    res := 0
    left := 0
    right := left + 1

    if len(nums) < 2 {
        if nums[0] == 1 {
            return 1
        } else {
            return 0
        }
    }

    for {
        // 不含1的数组
        if left == len(nums) - 1 {
            return 0
        }

        if right >= len(nums) - 1{
            break
        }

        // 如果起始位是1，则开始右移
        if nums[left] == 1 {

            if nums[right] == 1 {
				distance := right - left + 1
                right++
                if distance > res {
                    res = distance
                }else {
                    left = right
                    right = left + 1
                }
            }else {
				left = right + 1
            	right = left + 1
                
            }
        } else {
			left++
		}      
    }

    return res
}