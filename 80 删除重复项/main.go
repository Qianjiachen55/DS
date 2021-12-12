package main


func removeDuplicates(nums []int) int {
	// sorts.Int(nums)
	n:=len(nums)
	if len(nums)<=2{
		return n
	}
	i ,j :=2,2
	for j<n{
		if nums[i-2]!=nums[j]{
			nums[i]=nums[j]
			i++
		}
		j++
	}
	return i

}
