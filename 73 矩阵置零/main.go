package main


func setZeroes(matrix [][]int)  {
	m,n := len(matrix),len(matrix[0])
	mFlag := map[int]bool{}
	nFlag := map[int]bool{}

	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			if matrix[i][j]==0{
				nFlag[j] = true
				mFlag[i] = true
			}
		}
	}

	for k := range nFlag{
		for i:=0;i<m;i++{
			matrix[i][k]=0
		}
	}
	for k:= range mFlag{
		for j:=0;j<n;j++{
			matrix[k][j]=0
		}
	}

}
