package main

func mySqrt(x int) int {
	if x<=1{
		return x
	}
	left,right := 0,x
	mid := 0
	for left<right-1{
		mid = (left+right)/2
		temp := mid * mid
		if temp < x{
			left = mid
		}else if(temp==x){
			return mid
		}else{
			right = mid
		}
	}
	return left
}
