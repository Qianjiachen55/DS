package main

func solveSudoku(board [][]byte)  {
	backtrack(board)
}


func backtrack(board [][]byte)bool {
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]=='.'{
				var c byte
				for c='1';c<='9';c++{
					if isValid(board,i,j,c){
						board[i][j]=c

						if backtrack(board){
							return true
						}else{
							board[i][j]='.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		// check列，即每一行上的第col列是否为出现过c
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		// check行，即第row行上的每一列是否为出现过c
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		// check 3*3 block
		boxRow, boxCol := (row/3)*3+i/3, (col/3)*3+i%3
		if board[boxRow][boxCol] == c {
			return false
		}
	}
	return true
}
