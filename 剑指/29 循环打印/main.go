package main


func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0|| len(matrix[0])==0{
		return []int{}
	}
	up,down,left,right := 0,len(matrix)-1,0,len(matrix[0])-1
	res :=[]int{}
	for up <=down && left <=right{
		// ->
		for i:=left;i<=right;i++{
			res = append(res,matrix[up][i])
		}
		up++
		// xia
		for i:=up;i<=down;i++{
			res = append(res,matrix[i][right])
		}
		right--
		if left>right || up > down{
			break
		}
		//<-
		for i:=right;i>=left;i--{
			res = append(res,matrix[down][i])
		}
		down--
		//^
		for i:=down;i>=up;i-- {
			res = append(res,matrix[i][left])
		}
		left++
	}
	return res
}
