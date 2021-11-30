package main


func lengthOfLastWord(s string) int {
	n:=len(s)-1
	ans := 0

	for n>=0{
		if s[n]!=' '{
			break
		}
		n--
	}
	for n>=0{
		if s[n]==' '{
			break
		}
		n--
		ans++
	}
	return ans
}
