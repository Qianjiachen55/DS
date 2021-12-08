package main

func climbStairs(n int) int {
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
	if n==3{
		return 3
	}
	a := 1
	b := 2
	ans := 0
	for i:=3;i<=n;i++{
		ans = a+b
		a = b
		b = ans
	}

	return ans
}
