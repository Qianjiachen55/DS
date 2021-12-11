package main


func subsets(nums []int) [][]int {

	ans:=[][]int{}


	var dfs func(start int)
	cur := []int{}

	m := map[int]bool{}
	for i :=0;i<len(nums);i++{
		m[i]=false
	}
	dfs = func(start int){
		temp := make([]int,len(cur))
		copy(temp,cur)
		ans = append(ans,temp)

		for i:=start;i<len(nums);i++{
			if !m[i]{
				cur = append(cur,nums[i])
				m[i]=true
				dfs(i+1)

				cur = cur[:len(cur)-1]
				m[i]= false
			}
		}

		return
	}


	dfs(0)

	return ans




}
