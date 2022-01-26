package main

func firstMissingPositive2(nums []int) int {
	n := len(nums)

	//1. 将<=0 的数转换为n+1
	// -->> nums[i] >0
	for i:=0;i<n;i++{
		if nums[i]<=0{
			nums[i] = n+1
		}
	}
	// nums 全部为+ index:value  +(无)/ -(有)
	//n=4
	//3,4,5,1
	//-3,4,-5,-1
	// nums[i]< n+1 --> nums[num[i]] = -abs(value)
	for i:=0;i<n;i++{
		if abs2(nums[i])<n+1{
			index :=abs2(nums[i])-1
			nums[ index] = -abs2(nums[ index])
		}
	}

	for i:=0;i<n;i++{
		if nums[i]>0{
			return i+1
		}
	}

	return n+1


}


func abs2(n int)int{
	if n >=0{
		return n
	}
	return -n
}
