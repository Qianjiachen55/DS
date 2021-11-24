package main

func jump(nums []int) int {
	left,right := 0,0
	index := 0
	ans := 0
	rightTemp :=right
	for right < len(nums)-1{
		index = left
		ans +=1
		for index <=right{
			rightTemp = max(rightTemp,index+nums[index])
			index+=1
		}
		left = right+1
		right = rightTemp
	}
	return ans
}
func max(a ,b int)int{
	if a> b{
		return a
	}
	return b
}
