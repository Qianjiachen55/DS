package main

func totalNQueens(n int) int {
	board := make([][]bool,n)
	ans := 0
	for i:=0;i<n;i++{
		board[i] = make([]bool,n)
	}

	var backTrace func(board [][]bool,n int)
	backTrace =func(board [][]bool,n int){
		if n>=len(board){
			ans += 1
			return
		}
		for i:=0;i<len(board);i++{
			if isVaild(n,i,board){
				board[n][i]=true
				backTrace(board,n+1)
				board[n][i]=false
			}
		}

	}
	backTrace(board,0)

	return ans
}


func isVaild(x,y int ,board [][]bool)bool{
	for i:=0;i<x;i++{
		if board[i][y]{
			return false
		}
	}
	for i:=0;i<len(board);i++{
		if board[x][i]{
			return false
		}
	}
	for i,j := x,y;i>=0&&j>=0;i,j=i-1,j-1{
		if board[i][j]{
			return false
		}
	}
	for i,j:=x,y;i>=0&&j<len(board);i,j = i-1,j+1{
		if board[i][j]{
			return false
		}
	}




	return true
}
