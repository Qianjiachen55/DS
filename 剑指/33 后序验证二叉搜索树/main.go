package main


func verifyPostorder(postorder []int) bool {



	var helper func(start,end int)bool

	helper = func(start ,end int) bool{
		if start >=end{
			return true
		}
		p := start
		for postorder[p] < postorder[end]{
			p++
		}
		m := p
		for postorder[p]> postorder[end]{
			p++
		}

		return p==end && helper(start,m-1)&&helper(m,end-1)

	}

	return helper(0,len(postorder)-1)
}
