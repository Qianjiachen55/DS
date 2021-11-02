package main

import (
	"fmt"
)

func isValid(s string) bool {

	if len(s) % 2!=0{
		return false
	}
	stack := make([]byte,0)
	stack = append(stack,s[0])
	hash := map[byte]byte{')':'(',']':'[','}':'{'}
	for i := 1;i<len(s);i++{
		if s[i]=='(' || s[i]== '{' || s[i]=='['{
			stack = append(stack,s[i])
		} else if stack[len(stack)-1]== hash[s[i]]{
			stack = stack[:len(stack)-1]
		}else {
			return false
		}

	}
	if len(stack)==0{
		return true
	}
	return false

}


func main() {
	fmt.Println(isValid("()"))
}