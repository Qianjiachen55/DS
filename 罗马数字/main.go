package main

import "fmt"

func intToRoman(num int) string {
	hash :=map[int]string{
		1000:"M",
		900:"IM",
		500:"D",
		400:"ID",
		100:"C",
		90:"IC",
		50:"L",
		40:"IL",
		10:"X",
		9:"IX",
		5:"V",
		4:"IV",
		1:"I",
	}
	indexSlice := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	ans := ""
	i := 0
	for num != 0{
		strCount := num / indexSlice[i]
		num = num % indexSlice[i]
		for j := 0;j<strCount;j++{
			ans+=hash[indexSlice[i]]
		}
		i++
	}
	return ans
}

func main()  {
	res := intToRoman(9)
	fmt.Println("res:",res)
}