package leetcode

func MoveZeroes(nums []int)  {

	// 直接暴力破解，这样的方法执行时间比较久，n方
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == 0 {
            for j := i + 1; j < len(nums); j++ {
                // swap two elements
                if nums[j] != 0 && nums[i] == 0 {
                    nums[i], nums[j] = nums[j], nums[i]
                    i++
                }            
            }
        }
    }
}

func MoveZeroesPlus(nums []int)  {

    i := 0
    j := 0

	// 直接将不是0的元素先存在前面，后面的就继续补0
    for ; i < len(nums); i++ {
        if (nums[i] != 0) {
            nums[j] = nums[i]
            j++
        }
    }

    for {
        if j >= len(nums) {
            break
        }
        nums[j] = 0
        j++
    }
}