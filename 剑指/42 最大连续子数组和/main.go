package main


func maxSubArray(nums []int) int {
	if len(nums)==0{
		return 0
	}
	dp := make([]int,len(nums))
	dp [0] = nums[0]
	res := dp[0]

	for i:=1;i<len(nums);i++{
		if dp[i-1]<0{
			dp[i] = nums[i]
		}else{
			dp[i] = dp[i-1]+nums[i]
		}
		res = max(dp[i],res)
	}

	return res
}



func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}
