package main


func cuttingRope(n int) int {
	if n <4{
		return n-1
	}

	res := 1
	t := 1000000007
	for n > 4{
		res =res *3%t
		n -=3
	}
	// 4 3 2
	res = res*n%t

	return res
}
