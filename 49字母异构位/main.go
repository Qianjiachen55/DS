package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _,str:=range strs{
		temp := []byte(str)
		sort.Slice(temp,func(i,j int)bool{return temp[i]<temp[j]})
		sortedStr:=string(temp)
		m[sortedStr]=append(m[sortedStr],str)
	}
	ans := [][]string{}
	for _,v :=range m{
		ans = append(ans,v)
	}


	return  ans
}
