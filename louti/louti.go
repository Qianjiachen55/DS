package main

import "time"

func main() {
	t1 := time.Now().UnixNano()
	print(demo1(10),"\n")
	t2 := time.Now().UnixNano()
	print(t2-t1,"\n")
	t3 := time.Now().UnixNano()
	print(demo2(10),"\n")
	t4 := time.Now().UnixNano()

	print(t4-t3,"\n")
	//print(time.Now().UnixNano(),"\n")
	//print(time.Now().UnixNano())
}

func demo1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return demo1(n-1) + demo2(n-2)
}

func demo2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first := 1
	second := 2
	var ans int
	for i:=3;i<=n;i++{
		ans = first + second
		first = second
		second = ans
	}
	return ans

}
