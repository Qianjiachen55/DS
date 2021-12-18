package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int,0)
	n := len(nums)
	used := make([]bool,n)
	cur := []int{}
	curIndex :=0

	var dfs func()
	dfs = func(){
		if curIndex>=n{
			temp := make([]int,len(cur))
			copy(temp,cur)
			ans = append(ans,temp)
			return
		}

		if curIndex>0 && nums[curIndex]==nums[curIndex-1] && used[curIndex-1]==false{
			// 0 不能取
			curIndex++
			dfs()
			curIndex--
		}else{

			// 可取可不取
			//不取 0
			curIndex++
			dfs()
			curIndex--

			//取 1

			cur = append(cur,nums[curIndex])
			used[curIndex]=true
			curIndex++

			dfs()

			curIndex--
			used[curIndex] = false
			cur = cur[:len(cur)-1]


		}

	}
	dfs()

	return ans

}
