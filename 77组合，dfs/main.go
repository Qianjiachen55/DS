package main


func combine(n int, k int) [][]int {
	ans :=[][]int{}
	numsMap :=map[int]bool{}
	for i:=1;i<=n;i++{
		numsMap[i]=false
	}
	cur :=[]int{}

	var dfs func(start int)
	dfs =func(start int){
		if len(cur)==k{
			temp := make([]int,k)
			copy(temp,cur)
			ans = append(ans,temp)
			return
		}
		for i:=start;i<=n;i++{
			if !numsMap[i]{
				cur = append(cur,i)
				numsMap[i]=true
				dfs(i +1)

				cur = cur[:len(cur)-1]
				numsMap[i]=false
			}
		}

	}

	dfs(1)

	return ans
}
