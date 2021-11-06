package main

import "fmt"

func divide(dividend int, divisor int) int {
	flag := 1
	if dividend < 0 && divisor >0{
		dividend = 0-dividend
		flag =0-flag
	}
	if dividend > 0 && divisor <0{
		divisor = 0-divisor
		flag = 0-flag
	}
	if dividend <0 && divisor <0{
		dividend = 0-dividend
		divisor = 0-divisor
	}
	ans :=0
	_max :=(1<< 31) -1
	_min :=-( 1<<31)
	for dividend >= divisor{

		dividend = dividend-divisor
		ans +=1
		if ans > _max||ans < _min{
			return _max
		}
	}

	if flag==-1{
		ans = 0-ans
	}
	return ans
}

func main() {
	res := divide(-2147483648,-1)
	fmt.Println(res)
}