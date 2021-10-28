package main

import (
	"fmt"
)

func main()  {
	fmt.Println(myAtoi("18446744073709551617"))
}


func myAtoi(s string) int {
	ans := 0
	sTemp := s
	isZ := 1
	length := len(s)
	if len(s) == 0 {
		return ans
	}
	index := 0
	for i := 0; i < length; i++ {
		if s[i] != ' ' {
			index = i
			break
		}
	}
	sTemp = s[index:]
	length = len(sTemp)
	if length == 0 {
		return ans
	}
	deletFlag := false
	if sTemp[0] == '-' {
		isZ = -1

		deletFlag = true
	}
	if sTemp[0] == '+' {
		deletFlag = true
	}
	if deletFlag{
		sTemp = sTemp[1:]
	}

	length = len(sTemp)
	intSlice := make([]int,0)
	for i := 0; i < length; i++ {
		if isNum(string(sTemp[i])) {
			intSlice = append(intSlice, int(sTemp[i])-'0')
		} else {
			break
		}
	}
	for _, intTemp := range intSlice {
		ans *= 10
		ans += intTemp
		if ans >=(1 << 31) {
			if isZ==1 {
				return (1 << 31) -1
			}else{
				return -(1 << 31)
			}
		}

	}

	ans *= isZ
	if int(int32(ans))!= ans{
		if isZ==-1{
			return -(1 << 31)
		} else{
			return (1 << 31) -1
		}
	}


	return ans

}


func isNum(ch string) bool {
	switch ch {
	case "0":
		return true
	case "1":
		return true
	case "2":
		return true
	case "3":
		return true
	case "4":
		return true
	case "5":
		return true
	case "6":
		return true
	case "7":
		return true
	case "8":
		return true
	case "9":
		return true
	default:
		return false
	}
	return false

}
