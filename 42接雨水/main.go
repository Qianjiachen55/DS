package main

import "fmt"

func trap(height []int) int {
	n:=len(height)

	if n==0{
		return 0
	}
	leftMax := make([]int,n)
	leftMax[0] = height[0]
	for i:=1;i<n;i++{
		leftMax[i] = max(leftMax[i-1],height[i])
	}

	rightMax :=make([]int,n)
	rightMax[n-1] = height[n-1]
	for i:=n-2;i>=0;i--{
		rightMax[i] = max(rightMax[i+1],height[i])
	}
	ans := 0
	for i:=0;i<n;i++{

		ans += leftMax[i]+rightMax[i]-height[i]-leftMax[n-1]
		fmt.Println(leftMax[i]+rightMax[i]-height[i]-leftMax[n-1])
	}
	return ans

}
func max(a ,b int)int{
	if a >b{
		return a
	}
	return b
}

func min(a,b int)int{
	if a > b{
		return b
	}
	return a
}

func main() {
	res :=trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})
	fmt.Println(res)
}