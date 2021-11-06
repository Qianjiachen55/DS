package main

import "fmt"

func search(nums []int, target int) int {
	length := len(nums)
	if length == 1{
		if target == nums[0]{
			return 0
		}else{
			return -1
		}
	}
	l:=0
	r:=length-1
	mid := (l+r)/2
	for l<=r{

		mid = (l+r)/2
		fmt.Println(l,r,mid)
		if target == nums[mid]{
			return mid
		}
		if nums[l]<=nums[mid]{
			if nums[l]<=target && target<= nums[mid]{
				r = mid
			}else{
				l = mid+1
			}
		}else{
			if nums[mid+1]<=target && target <=nums[r]{
				l = mid+1
			}else{
				r = mid
			}
		}
	}
	return -1
}
func main()  {
	res := search([]int{1,3},3)
	fmt.Println(res)
}
