package main

func searchRange(nums []int, target int) []int {
	if len(nums)==0{
		return  []int{-1,-1}
	}else{
		return []int{searchLeft(nums,target),searchRight(nums,target)}
	}
}

func searchLeft(nums []int,target int)int{
	l,r := 0,len(nums)-1
	for l<=r{
		m:=(l+r)/2
		if nums[m]>=target{
			r = m-1
		}else{
			l = m+1
		}
	}
	if l<len(nums) && nums[l]==target{
		return l
	}
	return -1
}

func searchRight(nums []int,target int)int{
	l,r := 0,len(nums)-1
	for l<=r{
		m:=(l+r)/2
		if nums[m]<=target{
			l = m+1
		}else{
			r = m-1
		}
	}
	if r >=0 &&nums[r]==target{
		return r
	}
	return -1
}
