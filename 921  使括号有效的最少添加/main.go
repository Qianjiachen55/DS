package main



func minAddToMakeValid(s string) int {
	length := len(s)
	if length ==0 || length ==1{
		return length
	}
	stack := []byte{}

	for i:=0;i<length;i++{
		if s[i]=='('{
			stack = append(stack,s[i])
		}else{
			// s[i]=')'
			if len(stack) > 0 && stack[len(stack)-1]=='('{
				stack = stack[:len(stack)-1]
			}else{
				stack = append(stack,s[i])
			}
		}
	}
	return len(stack)
}