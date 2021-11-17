package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
	qsort(candidates,0,len(candidates)-1)
	fmt.Println(candidates)
	res := [][]int{}


	dfs(candidates,target,0,[]int{},&res)
	return res
}


func dfs(candidates []int,target int,index int,nums []int, res *[][]int){
	if target==0{
		temp := make([]int,len(nums))
		copy(temp,nums)
		*res = append(*res,temp)
		return
	}
	for i:=index;i<len(candidates);i++{
		if target < candidates[i]{
			return
		}
		dfs(candidates,target-candidates[i],i+1,append(nums,candidates[i]),res)
	}
}

func qsort(arr []int,start,end int){
	if start < end{
		index := partition(arr,start,end)
		qsort(arr,start,index-1)
		qsort(arr,index+1,end)
	}
}

func partition(arr []int,start,end int) int{
	l,r := start,end
	temp := arr[start]
	for l<r{
		for l<r && temp <=arr[r]{
			r--
		}
		for l<r && temp>=arr[l]{
			l++
		}
		arr[r],arr[l] = arr[l],arr[r]
	}
	arr[l],arr[start] = arr[start],arr[l]
	return r
}

func main()  {
	res := combinationSum2([]int{10,1,2,7,6,1,5,8},8)
	fmt.Println(res)
}