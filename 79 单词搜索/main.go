package main


func exist(board [][]byte, word string) bool {

	m :=len(board)
	n := len(board[0])

	flag :=make([][]bool,m)
	for i:=0;i<m;i++{
		flag[i] = make([]bool,n)
	}

	var dfs func(row int,col int)bool
	cur := ""
	curIndex :=0
	ops :=[][]int{
		{1,0},
		{0,1},
		{-1,0},
		{0,-1},
	}
	dfs = func(row int,col int)bool{
		if cur == word{
			return true
		}
		if len(cur)>=len(word){
			return false
		}
		if row >=0 && row <m && col >=0 && col <n && !flag[row][col] {
			if board[row][col]==word[curIndex]{
				cur += string([]byte{board[row][col]})
				curIndex++
				flag[row][col]=true
				for k:=0;k<len(ops);k++{
					if dfs(row+ops[k][0],col+ops[k][1]){
						return true
					}
				}
				curIndex--
				cur = cur[:curIndex]
				flag[row][col]=false
			}
		}
		return false

	}


	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if dfs(i,j){
				return true
			}
		}
	}

	return false
}
