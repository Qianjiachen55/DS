package main

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s)<4 && len(s)>12{
		return []string{}
	}
	sLen := len(s)
	ans := []string{}
	point := []int{0}
	start := 0
	count := 0
	var dfs func(curIndex int)

	dfs = func(curIndex int){
		if count == 3 &&isVaild(s[point[3]:]) {
			temp := ""
			for i:=0;i<len(point)-1;i++{
				temp += s[point[i]:point[i+1]]+"."
			}
			temp += s[point[3]:]
			ans = append(ans,temp)
			return
		}

		for i := curIndex+1;i<sLen;i++{
			if sLen-i > 3*(4-count-1){
				continue
			}

			if isVaild(s[point[count]:i]){
				point = append(point,i)
				count +=1
				dfs(i)
				point = point[:len(point)-1]
				count -=1
			}else{
				break
			}
		}
	}

	dfs(start)

	return ans


}


func isVaild(s string) bool{
	if len(s) > 3{
		return false
	}
	if len(s)<=0{
		return false
	}

	sInt,_ := strconv.Atoi(s)
	if s[0]=='0' && len(s)>1 {
		return false
	}

	if sInt>=0 && sInt<=255{
		return true
	}

	return false

}
