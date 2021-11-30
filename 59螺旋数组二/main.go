package main


func generateMatrix(n int) [][]int {
	up,down := 0,n-1
	left,right := 0,n-1
	ans := make([][]int,n)
	for index:=0;index<n;index++{
		ans[index] = make([]int,n)
	}
	i,j := 0,0
	sumN:=n*n
	for count:=1;count<=sumN;count++{
		ans[i][j]=count
		// 左上->右上
		for j<right&& count<sumN{
			count++
			j++
			ans[i][j]=count
		}
		up++
		for i<down&& count<sumN{
			count ++
			i++
			ans[i][j] =count
		}
		right--
		for j>left&& count<sumN{
			count ++
			j--
			ans[i][j]=count
		}
		down--
		for i>up&& count<sumN{
			count ++
			i--
			ans[i][j]=count
		}
		left++
		j++

	}
	return ans

}
