package main

import (
	"fmt"
)

func main()  {
	m := map[string][]string{
		"1":{"a","b","c"},
		"2":{"d","e"},
		"3":{"f","g","h"},
	}

	str := "223"

	res := []string{}

	cur :=""

	var dfs func()
	dfs = func() {
		if len(str)==0{
			var temp string
			temp = cur
			res = append(res,temp)
			return
		}
		tempS := string(str[0])
		choiceS := m[tempS]
		for i:=0;i<len(choiceS);i++{
			cur = cur+choiceS[i]
			str = str[1:]
			dfs()
			str = tempS+str
			cur = cur[:len(cur)-1]
		}

		return
	}


	dfs()


	for _,v:=range res{
		fmt.Println(v)
	}

	fmt.Println("done!!!!")
}
