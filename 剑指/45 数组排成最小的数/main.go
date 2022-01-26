package main

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	compare := func(i,j int)bool{
		a,b := 0,0

		tempI,tempJ := nums[i],nums[j]


		exp := 1
		for tempI >0{
			a = a + (tempI%10*exp)
			tempI /=10
			exp *=10
		}
		if tempJ == 0{
			a *=10
		}
		for tempJ>0{
			tempJ/=10
			a *=10
		}
		a = a+nums[j]

		tempI,tempJ = nums[i],nums[j]
		exp = 1
		for tempJ >0 {
			b = b+(tempJ%10*exp)
			tempJ /=10
			exp *=10
		}
		if tempI ==0{
			b *=10
		}
		for tempI >0{
			tempI /=10
			b*=10
		}
		b = b+nums[i]
		return a<b

	}

	sort.Slice(nums, compare)

	res := ""
	for i:=0;i<len(nums);i++{

		res += strconv.Itoa(nums[i])
	}
	if len(res)==0{
		return "0"
	}
	// if len(res)==1{
	//     return res
	// }
	// flag :=0
	// for ;flag<len(res)-1&&res[flag]=='0';flag++{
	// }

	return res

}
