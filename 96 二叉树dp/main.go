package main


func numTrees(n int) int {
	ans :=make([]int,n+1)
	ans[0],ans[1] = 1,1

	for i:=2;i<=n;i++{
		for j:=0;j<i;j++{
			ans[i] += ans[j] * ans[i-j-1]
		}
	}

	return ans[n]
}
