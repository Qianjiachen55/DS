package main

import "fmt"

func numDecodings(s string) int {

	if len(s)==0{
		return 0
	}
	if s[0]=='0'{
		return 0
	}

	l := len(s)+2
	dp := make([]int,l)
	dp[0],dp[1] = 1,1
	dp[2] = 1

	for i:=3;i < l;i++{
		if s[i-2]=='0'&&(s[i-3]=='1'||s[i-3]=='2'){
			dp[i] = dp[i-2]
		}

		if s[i-3]=='1' && s[i-2]>='1' && s[i-2]<='9'{
			dp[i] = dp[i-1]+dp[i-2]
		}

		if s[i-3]=='2' && s[i-2]>='1'&&s[i-2]<='6'{
			dp[i] = dp[i-1]+dp[i-2]
		}

		if s[i-3]=='0'&& s[i-2]>='1'{
			dp[i] = dp[i-1]
		}

		if s[i-3]=='2' && s[i-2]>='7'{
			dp[i] = dp[i-1]
		}
		if s[i-3]>='3' && s[i-2]>'0'{
			dp[i] = dp[i-1]
		}


	}
	fmt.Println(dp)

	return dp[l-1]
}