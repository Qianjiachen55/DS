package main



func plusOne(digits []int) []int {
	add :=1
	for i:=len(digits)-1;i>=0;i--{
		if digits[i]==9 && add ==1{
			digits[i]=0
			add = 1
		}else{
			digits[i]+=add
			add = 0
			break
		}
	}
	if add ==1{
		return append([]int{1},digits...)
	}
	return digits

}
