package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	right, n := 0,len(words)
	ans :=[]string{}
	for {
		left := right

		sumLen := 0
		for right < n && sumLen+len(words[right])+right-left <= maxWidth{
			sumLen += len(words[right])
			right++
		}

		if right == n{
			s := strings.Join(words[left:]," ")
			ans = append(ans,s+blank(maxWidth-len(s)))
			return ans
		}

		numWords := right - left
		numSpace := maxWidth - sumLen

		if numWords ==1{
			ans = append(ans, words[left]+blank(numSpace))
			continue
		}
		// else
		avgSpaces := numSpace/(numWords - 1)
		extraSpaces := numSpace%(numWords-1)
		s1 := strings.Join(words[left:left+extraSpaces+1],blank(avgSpaces+1))
		s2 := strings.Join(words[left+extraSpaces+1:right],blank(avgSpaces))
		ans = append(ans,s1+blank(avgSpaces)+s2)
	}
	return ans
}

func blank (n int)string{
	return strings.Repeat(" ",n)
}
