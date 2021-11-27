package main

import "strings"

func solveNQueens(n int) [][]string {
	res := [][]string{}
	board := make([][]string,n)
	for i:=0;i<n;i++{
		board[i] = make([]string,n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			board[i][j] ="."
		}
	}
	var backTrace func(board [][]string,n int)
	backTrace = func(board [][]string,n int){
		if n>=len(board){
			temp := make([]string,len(board))
			for i:=0;i<len(board);i++{
				temp[i] =strings.Join(board[i],"")
			}
			res = append(res,temp)
			return
		}
		for i:=0;i<len(board);i++{
			if isVaild(n,i,board){
				board[n][i]="Q"
				backTrace(board,n+1)
				board[n][i]="."
			}
		}

		return

	}
	backTrace(board,0)

	return res

}


func isVaild(x, y int,board [][]string)(bool){
	n := len(board)

	for i:=0;i<x;i++{
		if board[i][y]=="Q"{
			return false
		}
	}
	for i:=0;i<n;i++{
		if board[x][i]=="Q" && i!=y{
			return false
		}
	}
	for i,j:=x,y;i>=0 && j>=0;i,j =i-1,j-1{
		if board[i][j]=="Q"{
			return false
		}
	}
	for i,j:=x,y;i>=0&&j<n;i,j=i-1,j+1{
		if board[i][j]=="Q"{
			return false
		}
	}

	return true


}