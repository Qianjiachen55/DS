package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i,j int)bool{
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	start := intervals[0][0]
	end := intervals[0][1]
	for i:=1;i<len(intervals);i++{
		if end >= intervals[i][0]{
			end = max(intervals[i][1],end)
		}else{
			res = append(res,[]int{start,end})
			start ,end = intervals[i][0],intervals[i][1]
		}
	}
	res = append(res,[]int{start,end})
	return res
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}
