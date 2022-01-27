package main


func numIslands(grid [][]byte) int {

	if len(grid)==0||len(grid[0])==0{
		return 0
	}

	m ,n := len(grid),len(grid[0])
	ans :=0
	visited := make([][]bool,m)
	for i:=0;i<m;i++{
		visited[i] = make([]bool,n)
	}
	target :=byte('1')

	bfs :=func(i,j int)bool{


		//queue queue[i]==1
		queue := [][]int{}
		visited[i][j] = true
		if grid[i][j]!=target{
			return false
		}

		queue = append(queue,[]int{i,j})
		var x,y int
		for len(queue)>0{
			x,y = queue[0][0],queue[0][1]

			queue = queue[1:]


			// up
			if x-1 >=0 && !visited[x-1][y] && grid[x-1][y]==target{
				queue = append(queue,[]int{x-1,y})
				visited[x-1][y] = true
			}
			// down
			if x+1<m && !visited[x+1][y] && grid[x+1][y]==target{
				queue = append(queue,[]int{x+1,y})
				visited[x+1][y] = true
			}
			//left
			if y-1 >=0 && !visited[x][y-1] && grid[x][y-1]==target{
				queue = append(queue,[]int{x,y-1})
				visited[x][y-1]=true
			}

			//right
			if y+1<n && !visited[x][y+1] && grid[x][y+1]==target{
				queue = append(queue,[]int{x,y+1})
				visited[x][y+1]=true
			}


		}

		return true


	}


	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if !visited[i][j]&&bfs(i,j){
				ans ++
			}
		}
	}


	return ans


}
