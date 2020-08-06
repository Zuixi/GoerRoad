package leetcode

type Employee struct {
	Id int
	Importance int
	Subordinates []int
}


// #690
func GetImportance(employees []*Employee, id int) int {
    
    m := map[int]*Employee{}
	for _, v := range employees {
		m[v.Id] = v
	}

	res := 0
	
	// 这里使用匿名方法写dfs
	var dfs func(i int)
	dfs = func(i int) {
		res += m[i].Importance
		for _, v := range m[i].Subordinates {
			dfs(v)
		}
	}

	dfs(id)
	return res  
}