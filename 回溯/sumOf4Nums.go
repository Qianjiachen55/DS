package main

import (
	"fmt"
	"sort"
)

var ans [][]int
var subAns,myNums [] int
var ansHash map[int]struct{}
var tar,numSize int

func DFS(low int,sum int){
	if sum == tar && len(subAns)==4{
		temp := make([]int,0)
		temp = append(temp,subAns...)
		//copy(temp,subAns)
		ans = append(ans, temp)
		fmt.Println(temp)
		//fmt.Println(subAns)
		fmt.Println(ans)
		return
	}
	//fmt.Println(subAns)
	for i:=low;i<numSize;i++{

		//if numSize - i < 4-len(subAns){
		//	return
		//}
		//if i > low && myNums[i-1]==myNums[i]{
		//	continue
		//}
		//if i < numSize -1 && sum
		subAns = append(subAns,myNums[i])
		DFS(i+1,myNums[i]+sum)
		//fmt.Println(i,subAns)
		subAns = subAns[:len(subAns)-1]
		//fmt.Println(i,subAns)
	}
	//fmt.Println(subAns)


}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	myNums = nums
	ansHash = make(map[int]struct{},0)
	tar = target
	numSize = len(nums)

	DFS(0,0)
	fmt.Println("=======")
	fmt.Println(ans)
	fmt.Println("=======")
	return ans
}

func main() {
	ans = fourSum([]int{1,0,-1,0,-2,2},0)
	//ans = fourSum([]int{-2,-1,0},0)
	fmt.Println(ans)
	fmt.Println(subAns)
}
