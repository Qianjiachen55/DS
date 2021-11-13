package main
func isValidSudoku(board [][]byte) bool {
	//0,0 0,3 0,6
	//3,0 3,3......
	for x:=0;x<9;x++{
		for y:=0;y<9;y++{
			if x%3==0 && y%3==0{
				if !isValid3c3(board,x,y){
					return false
				}
			}
			if !isValidXY(board,x,y){
				return false
			}
		}
	}
	return true

}


func isValidXY(board [][]byte,x,y int)bool{
	//判断xy所在行列是否合法
	r1 := map[byte]bool{}
	for i:=0;i<9;i++{
		if board[x][i]=='.'{
			continue
		}
		if r1[board[x][i]]{
			return false
		}
		r1[board[x][i]]=true
	}

	r2 := map[byte]bool{}
	for i:=0;i<9;i++{
		if board[i][y]=='.'{
			continue
		}
		if r2[board[i][y]]{
			return false
		}
		r2[board[i][y]]=true
	}
	return true

}

func isValid3c3(board [][]byte,x,y int)bool{
	v := map[byte]bool{}

	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if board[x+i][y+j]=='.'{
				continue
			}
			if _,ok := v[board[x+i][y+j]];ok{
				return false
			}
			v[board[x+i][y+j]]=true
		}
	}
	return true
}
