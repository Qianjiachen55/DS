package main

//利用自身数组做哈希表 index：value
func firstMissingPositive(nums []int) int {
	n:=len(nums)
	for i:=0;i<n;i++{
		if nums[i]<=0{
			nums[i] = n+1
		}
	}
	for i:=0;i<n;i++{
		num := abs(nums[i])
		if num <=n{
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i:=0;i<n;i++{
		if nums[i]>0{
			return i+1
		}
	}

	return n+1

}

func abs(num int)int{
	if num < 0{
		return -num
	}
	return num
}
