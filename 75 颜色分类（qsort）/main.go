package main


func sortColors(nums []int)  {
	qSort(nums,0,len(nums)-1)
}

func qSort(nums []int,start int,end int){
	if start > end{
		return
	}
	index := partition(nums,start,end)
	qSort(nums,start,index-1)
	qSort(nums,index+1,end)
}


func partition(nums []int,start int,end int) int {
	l,r := start, end
	temp := nums[l]
	for l < r{
		for l< r && nums[r] >= temp{
			r --
		}

		for l<r && nums[l] <= temp{
			l++
		}
		nums[r],nums[l] = nums[l],nums[r]
	}
	nums[start],nums[r] = nums[r],nums[start]
	return r
}
