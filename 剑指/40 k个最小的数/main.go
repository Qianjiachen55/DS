package main


func getLeastNumbers(arr []int, k int) []int {
	start,end := 0,len(arr)-1
	temp := 0
	if k ==0 {
		return []int{}
	}
	for {
		temp = partition(arr,start,end)
		if temp==k-1{
			return arr[0:temp+1]
		}else if temp >k-1{
			//子区间长
			//1,2,3
			//4,3,2,1,5
			end = temp-1
		}else{
			// 子区间短
			// 1,2,3,4,5
			// 2,1, 3  5,4,9,7,10
			start = temp+1
		}

	}
	return []int{}
}


func partition(arr []int,start,end int) int{
	l,r := start,end
	p := arr[start]
	for l < r{
		for l <r && arr[r] >=p{
			r--
		}
		for l<r && arr[l] <=p{
			l++
		}
		arr[l],arr[r] = arr[r],arr[l]
	}

	arr[start],arr[l] = arr[l],arr[start]
	return l
}
