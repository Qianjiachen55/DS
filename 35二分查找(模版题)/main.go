package main


func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n==0{
		return 0
	}
	l,r := 0,n-1
	for l<=r{
		mid := (l+r)/2
		if target<nums[mid]{
			r = mid-1
		}else if target==nums[mid]{
			return mid
		}else{
			l = mid+1
		}
	}
	return l
}
