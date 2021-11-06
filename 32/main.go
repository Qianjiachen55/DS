package main

import "fmt"

func nextPermutation(nums []int)  {
	length := len(nums)
	if length ==0||length==1{
		return
	}
	index := length - 2
	for index >=0 {
		if nums[index]>= nums[index+1]{
			index--
		}else{
			break
		}
	}

	if index >=0{
		temp := nums[index]
		for j:=length-1;j>index;j--{
			if nums[j]>temp{
				nums[index],nums[j] = nums[j],nums[index]
				break
			}
		}
	}


	fmt.Println(index)
	reverse(nums[index+1:])
}

func reverse(a []int){
	i:=0
	j:=len(a)-1
	for i<j{
		a[i],a[j] = a[j],a[i]
		i++
		j--
	}
}
