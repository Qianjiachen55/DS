package main


func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals)==0{
		return [][]int{newInterval}
	}
	tempI:=0
	i :=0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		i++
	}
	tempI = i
	for i=tempI;i<len(intervals)&&intervals[i][0]<=newInterval[1];i++{
		newInterval[0] = min(newInterval[0],intervals[i][0])
		newInterval[1] = max(newInterval[1],intervals[i][1])
	}
	res := append(intervals[:tempI],append([][]int{newInterval},intervals[i:]...)...)

	return res


}


func max(a,b int) int{
	if a >b {
		return a
	}
	return b
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}
