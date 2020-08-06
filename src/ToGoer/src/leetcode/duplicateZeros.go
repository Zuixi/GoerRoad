package leetcode

func DuplicateZeros(arr []int)  {

    for i := 0; i < len(arr); i++ {
        if arr[i] == 0 {
            // arr[i + 1] == 0 and i + 1 -> i + 2
            for j := len(arr) - 1; j > i + 1; j-- {
                arr[j] = arr[j - 1]
            }
            if i + 1 < len(arr){
                arr[i + 1] = 0
			}    
			i++        
        }
    }
}