package main


func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
		if i==0{
			dp[i][0] = grid[i][0]
			for j:=1;j<n;j++{
				dp[i][j] = grid[i][j]+dp[i][j-1]
			}
		}else{
			dp[i][0] = grid[i][0]+dp[i-1][0]
			for j:=1;j<n;j++{
				dp[i][j] = min(dp[i-1][j],dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]


}



func min(a,b int)int{
	if a > b{
		return b
	}
	return a
}
