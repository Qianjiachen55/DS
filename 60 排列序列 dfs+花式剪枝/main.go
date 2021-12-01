package main

import "strconv"

func getPermutation(n int, k int) string {
	num := []string{}
	for i:=1;i<n+1;i++{
		num = append(num,strconv.Itoa(i))
	}
	res :=""
	var dfs func()()
	dfs = func(){
		if len(num)==2{
			if k==1{
				res += (num[0]+num[1])
				return
			}
			res += (num[1]+num[0])
			return
		}

		n:=len(num)
		base := jiecheng(n-1)
		for i:=0;i<n;i++{
			if k>base{
				k-=base
			}else{
				res += num[i]
				num = append(num[:i],num[i+1:]...)
				dfs()
				return
			}
		}
	}
	dfs()
	return res
}

func jiecheng(n int)int{
	res :=1
	for i:=1;i<=n;i++{
		res *= i
	}
	return res
}
