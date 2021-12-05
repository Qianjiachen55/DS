package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	m := len(obstacleGrid)
	dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	for i:=0;i<n;i++{
		if obstacleGrid[0][i]==0{
			dp[0][i] = 1
		}else{
			break
		}
	}

	for i:=0;i<m;i++{
		if obstacleGrid[i][0]==0{
			dp[i][0] = 1
		}else{
			break
		}
	}

	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			if obstacleGrid[i][j]==0{
				dp[i][j]=dp[i-1][j]+dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
