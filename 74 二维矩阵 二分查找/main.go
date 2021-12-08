package main

func searchMatrix(matrix [][]int, target int) bool {
	l,r := 0,len(matrix)-1
	for l<=r{
		mid := (l+r)/2
		if matrix[mid][0]<target{
			l = mid + 1
		}else if matrix[mid][0]==target{
			return true
		}else{
			r = mid - 1
		}
	}
	row := (l+r)/2
	left,right := 0,len(matrix[row])-1
	for left <= right {
		mid := (left+right)
		if matrix[row][mid]<target{
			left = mid +1
		}else if matrix[row][mid]== target{
			return true
		}else{
			right = mid-1
		}

	}
	return false
}
