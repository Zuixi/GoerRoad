package leetcode

func DetectCapitalUse(word string) bool {

    // case 全部大写
    // case 只有首字母大写
    // case 全部小写

    // 从第二个字母开始是否全部是小写，如果是全是则return ture， 否则 则考虑是否全部大写，不然直接return false

	length := len(word)

	// 特殊case mL mmm
	if (word[0] <= 'z' && word[0] >= 'a') && length != 1{
        // 首字母小写，这个时候只有全部小写才是true，其他为false

        j := 1
        f := true
        // 全部小写
        for ; j < length; j++ {
            if word[j] >= 'a' && word[j] <= 'z' {
                f = true
            } else {
                f = false
                break
            }
        }

        if f == false {
            return false
        }
        return true
	}
	
    var flag bool = true

    i := 1
    // 全部小写
    for ; i < length; i++ {
        if word[i] >= 'a' && word[i] <= 'z' {
            flag = true
        } else {
            flag = false
            break
        }
    }

    if flag == false {
         // 全部大写
        for i = 1; i < length; i++ {
            if word[i] >= 'A' && word[i] <= 'Z' {
                flag = true
            } else {
                flag = false
                break
            }
        }

        if flag == false {
            return false
        }
    }

    return true

}