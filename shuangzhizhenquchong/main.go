package main

//双指针去重

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length==0{
		return 0
	}
	low:=0
	fast:=0
	for fast <length{
		if nums[fast]==nums[low]{
			fast ++
		}else{
			low++
			nums[low]=nums[fast]
		}
	}
	return low+1
}
