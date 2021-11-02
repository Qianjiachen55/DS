package main

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(lRemain int,rRemain int,path string)
	dfs = func(lRemain int,rRemain int,path string){
		if n*2 == len(path){
			ans = append(ans,path)
		}
		if lRemain > 0{
			dfs(lRemain-1,rRemain,path+"(")
		}
		if lRemain < rRemain{
			dfs(lRemain,rRemain-1,path+")")
		}
	}

	dfs(n,n,"")
	return ans
}
