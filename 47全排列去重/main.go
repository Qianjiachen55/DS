package main

import "sort"

func permuteUnique(nums []int) [][]int {

	ans := [][]int{}
	sort.Ints(nums)
	used := make([]bool,len(nums))

	var backTrack func(nums []int,cur []int)()
	backTrack = func(nums [] int,cur []int){
		if len(cur)==len(nums){
			temp :=make([]int,len(nums))
			copy(temp,cur)
			ans = append(ans,temp)
			return
		}

		for i:=0;i<len(nums);i++{

			if i>0&&nums[i]==nums[i-1]&&used[i-1]==false{
				continue
			}

			if used[i]==false{
				used[i]=true
				cur = append(cur,nums[i])
				backTrack(nums,cur)
				used[i]=false
				cur = cur[:len(cur)-1]
			}
		}
	}
	backTrack(nums,[]int{})
	return ans
}
