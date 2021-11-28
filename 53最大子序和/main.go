package main


func maxSubArray(nums []int) int {
	preSum,maxSum := nums[0],nums[0]

	for i:=1;i<len(nums);i++{
		if preSum >=0{
			preSum += nums[i]
		}else{
			preSum = nums[i]
		}

		if preSum >maxSum{
			maxSum = preSum
		}
	}
	return maxSum
}
