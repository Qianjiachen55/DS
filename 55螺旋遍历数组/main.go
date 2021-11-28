package main

func spiralOrder(matrix [][]int) []int {
	up,down,left,right := 0,len(matrix)-1,0,len(matrix[0])-1
	res := []int{}
	i:=up
	j:=left
	size := (down+1)*(right+1)
	for up<=down && left<=right{
		res = append(res,matrix[i][j])
		for j<right && len(res)<size{
			j++
			res = append(res,matrix[i][j])
		}
		up++
		for i<down && len(res)<size{
			i++
			res = append(res,matrix[i][j])
		}
		right--
		for j>left && len(res)<size{
			j--
			res = append(res,matrix[i][j])
		}
		left ++
		// res = append(res,matrix[i][j])
		for i>up && len(res)<size{
			i--
			res = append(res,matrix[i][j])
		}

		down--
		j++

	}

	return res

}
