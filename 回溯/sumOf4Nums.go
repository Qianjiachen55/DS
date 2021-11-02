package main

import (
	"fmt"
	"sort"
)

var ans [][]int
var subAns,myNums [] int
var ansHash map[string]struct{}
var tar,numSize int



func DFS(low int,sum int){
	if sum == tar && len(subAns)==4{
		temp := make([]int,4)
		copy(temp,subAns)
		tempStr := toString(subAns)

		if _,ok:=ansHash[tempStr];!ok{
			ans = append(ans, temp)
			ansHash[tempStr] = struct{}{}
		}
		return
	}
	for i:=low;i<numSize;i++{
		if numSize-i <4-len(subAns){
			return
		}

		subAns = append(subAns,myNums[i])
		DFS(i+1,myNums[i]+sum)
		subAns = subAns[:len(subAns)-1]
	}


}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	myNums = nums
	ansHash = make(map[string]struct{})
	tar = target
	numSize = len(nums)

	DFS(0,0)
	//fmt.Println("=======")
	//fmt.Println(ans)
	//fmt.Println("=======")
	return ans
}

func toString(l []int) string {
	res := ""
	for _,i:=range l{
		res += fmt.Sprintf("%d",i)
	}
	//fmt.Println("str:",res)
	return res
}


func main() {
	ans = fourSum([]int{2,2,2,2,2},8)
	//ans = fourSum([]int{-2,-1,0},0)
	fmt.Println(ans)
	//fmt.Println(subAns)
}
