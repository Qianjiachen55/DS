package main


func main()  {

}


func isMatch(s string, p string) bool{


	//if len(p)==0{
	//	if len(s)==0{
	//		return true
	//	}else {
	//		return false
	//	}
	//}

	m := len(s) + 1
	n := len(p) + 1

	dp := make([][]bool,m)

	//dp[i][j]  p[:j] match s[:i]



	for i:=0;i<m;i++{
		dp[i] = make([]bool,n)
	}

	//初始化
	dp[0][0]=true
	for j:=2;j<n;j++{
		if p[j] == '*'{
			dp[0][j] = dp[0][j-2]
		}
	}


	for r:=1;r<m;r++{
		i := r-1
		for c:=1;c<n;c++{
			j:= c-1
			if s[i]==p[j] || p[j]=='.'{
				dp[r][c] = dp[r-1][c-1]
			}
			if p[j] == '*'{
				if p[j-1]=='.'{
					dp[r][c] = dp[r][c-2]
				}
				if p[j-1] ==s[i]{
					dp[r][c] = dp[r-1][c]
				}
			}


		}
	}
	return dp[m-1][n-1]


}