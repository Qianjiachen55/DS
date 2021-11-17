package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	qsort(candidates,0,len(candidates)-1)
	//sort.Ints(candidates)
	fmt.Println(candidates)
	res :=[][]int{}

	dfs(candidates,target,[]int{},0,&res)
	return res

}


func dfs(candidates []int,target int,nums []int,start int,res *[][]int){
	if target ==0{
		temp := make([]int,len(nums))
		copy(temp,nums)
		*res = append(*res,temp)
		return
	}

	for i:=start;i<len(candidates);i++{
		if target <candidates[i]{
			return
		}
		dfs(candidates,target-candidates[i],append(nums,candidates[i]),i+1,res)
	}
}

func qsort(arr []int,start,end int){
	l,r := start,end
	if l<r{
		index := partition(arr,start,end)
		qsort(arr,start,index-1)
		qsort(arr,index+1,end)
	}
	return
}
func partition(arr []int,start,end int)int{
	l,r := start,end
	temp := arr[start]
	for l<r{
		for l<r && temp <= arr[r]{
			r--
		}
		for l<r &&temp>=arr[l]{
			l++
		}

		arr[l],arr[r] = arr[r],arr[l]
	}
	arr[start],arr[l] = arr[l],arr[start]
	return r
}

func main() {
	res := combinationSum([]int{2,3,6,7},7)
	fmt.Println(res)
}
