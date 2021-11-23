package main

import "fmt"

func isMatch(s string, p string) bool {
	row ,column:= len(p),len(s)
	dp := make([][]bool,row+1)
	for i:=0;i<=row;i++{
		dp[i] = make([]bool,column+1)
	}
	dp[0][0] = true

	if row>0  {
		i := 1
		for i<=row{
			if p[i-1]!='*'{
				break
			}
			dp[i][0]=true
			i++
		}

	}
	for i:=1;i <= row;i++{

		for j:=1;j<= column;j++{
			if p[i-1]=='?' || p[i-1]==s[j-1]{
				dp[i][j] = dp[i-1][j-1]
			}
			if p[i-1] == '*'{
				dp[i][j] = (dp[i-1][j]||dp[i][j-1]||dp[i-1][j-1])
			}
		}
	}
	return dp[row][column]
}

func main() {
	res := isMatch("ho", "**ho")
	fmt.Println(res)
}
