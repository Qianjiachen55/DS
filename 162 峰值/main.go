package main

import "math"

func findPeakElement(nums []int) int {
	get := func(i int)int{
		if i >=len(nums) || i<0{
			return math.MinInt64
		}
		return nums[i]
	}
	left ,rigth := 0,len(nums)
	for left <rigth{
		mid := (left+rigth)/2
		if get(mid)>get(mid-1) && get(mid)>get(mid+1){
			return mid
		}
		if get(mid) <get(mid-1){
			rigth = mid
		}else{
			left = mid
		}
	}

	return -1
}
