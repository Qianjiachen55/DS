package main

import "container/list"

func movingCount(m int, n int, k int) int {
	queue := list.List{}
	visited := make([][]bool,m)
	for i :=0;i<m;i++{
		visited[i] = make([]bool,n)
	}
	x,y := 0,0
	queue.PushBack([]int{x,y})

	res := 0
	for queue.Len()>0{
		back := queue.Back()
		queue.Remove(back)
		bfs := back.Value.([]int)
		x = bfs[0]
		y = bfs[1]
		if x<0 || x>=m || y<0 || y>= n || visited[x][y] ||!isVaild(x,y,k){
			continue
		}else{
			visited[x][y]=true
			res ++
			queue.PushBack([]int{x,y+1})
			queue.PushBack([]int{x,y-1})
			queue.PushBack([]int{x-1,y})
			queue.PushBack([]int{x+1,y})

		}

	}

	return res
}



func isVaild(i,j,k int)bool{
	// var res bool
	temp := 0
	for i >0{
		temp += i%10
		i/=10
	}
	for j>0{
		temp += j%10
		j/=10
	}
	if temp >k {
		return false
	}
	return true
}
