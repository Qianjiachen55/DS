package main

func canJump(nums []int) bool {
	if len(nums)==1{
		return true
	}
	maxIndex :=0
	for i:=0;i<=maxIndex;i++{
		maxIndex = max(maxIndex,i+nums[i])
		if maxIndex >=len(nums)-1{
			return true
		}
	}
	return false
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}
