package main

import "fmt"

func main()  {
	ss:=longestPalindrome("aaaaa")
	fmt.Println(ss)
}


func longestPalindrome(s string) string {
	length:=len(s)
	if length < 2{
		return s
	}
	maxLen := 1
	maxLenStr := ""
	maxLenStr = s[0:1]


	dp:=make([][]bool,length)
	for i:=0;i<length;i++{
		dp[i] = make([]bool,length)
	}

	for i:=0;i<length;i++{
		dp[i][i] = true
		if i+1 <length && s[i]==s[i+1]{
			dp[i][i+1] =true
			maxLen = 2
			maxLenStr = s[i:i+2]
		}
	}

	for l:=3;l<=length;l++{
		for i:=0;i+l-1<length;i++{
			if dp[i+1][i+l-1-1] && (s[i]==s[i+l-1]){
				dp[i][i+l-1]=true
				if maxLen < l{
					maxLen = l
					maxLenStr = s[i:i+l]
				}

			}
		}
	}


	return maxLenStr

}


func max(x int,y int)int{
	if x > y{
		return x
	}
	return y
}