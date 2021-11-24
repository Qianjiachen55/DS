package main

func permute(nums []int) [][]int {
	ans := [][]int{}
	var quanpailie func(nums []int,temp []int)()
	quanpailie = func(nums []int,temp []int){

		if len(temp) ==len(nums){
			cur := make([]int,len(temp))
			copy(cur,temp)
			ans = append(ans,cur)
			return
		}
		for i:=0;i<len(nums);i++{
			exist := false
			for j:=0;j<len(temp);j++{
				if nums[i]==temp[j]{
					exist = true
					break
				}

			}
			if !exist{
				temp = append(temp,nums[i])
				quanpailie(nums,temp)
				temp = temp[:len(temp)-1]
			}
		}
	}
	t := []int{}
	quanpailie(nums,t)

	return ans
}
