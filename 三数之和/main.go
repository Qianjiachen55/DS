package main

import (
	"fmt"
	"sort"
)

func main()  {
	ans := threeSum([]int{0,0,0})
	fmt.Println(ans)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	print(nums)
	ans := make([][]int,0)
	for k:=0;k<len(nums)-2;k++{
		if nums[k]>0{
			return ans
		}
		if  k>0&&nums[k]==nums[k-1]{
			continue
		}
		i:=k+1
		j:=len(nums)-1

		for i<j{

			temp := nums[k]+nums[j]+nums[i]
			if temp==0{
				s := []int{nums[k],nums[i],nums[j]}
				ans = append(ans,s)
				j--
				i++
				for i<j&&nums[i]==nums[i-1]{
					i++
				}
				for i<j && nums[j]==nums[j+1]{
					j--
				}
			}
			if temp >0{
				j--
				for j>i && nums[j]==nums[j+1] {
					j--
				}
			}
			if temp <0{
				i++
				for i<j&&nums[i]==nums[i-1]{
					i++
				}
			}
		}

	}
	return ans
}
