package main

import (
	"fmt"
)

func main()  {
	fmt.Println(convert("asdfghjkl",3))
}

func convert(s string, numRows int) string {
	if numRows == 1{
		return s
	}
	res := ""
	resSlice := make([]string,numRows)
	n := numRows*2 - 2
	for i:=0;i<len(s);i++{
		x := i % n
		resSlice[min(x,n-x)] +=string(s[i])
	}
	for _,i :=range resSlice{
		res += i
	}
	return res
}

func  min(x int ,y int) int {
	if x > y{
		return y
	}
	return x
}


