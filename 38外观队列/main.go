package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	words :="1"
	for i:=1;i<n;i++{
		words = say(words)
		//fmt.Println(words)
	}
	return words
}

func say(words string) string {
	res :=bytes.Buffer{}
	for i:=0;i<len(words);i++{
		count := 1
		for i+1<len(words) && words[i+1]==words[i]{
			count++
			i++
		}

		res.WriteString(strconv.Itoa(count))
		res.WriteByte(words[i])

	}
	return res.String()
}

func main() {
	fmt.Println(countAndSay(3))
}