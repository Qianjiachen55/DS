package main

func strStr(haystack string, needle string) int {

	if len(needle)==0{
		return 0
	}

	for i:=0;i+len(needle)-1<len(haystack);i++{
		for j:=0;j<len(needle);j++{
			if haystack[i+j]==needle[j] {
				if j==len(needle)-1{
					return i
				}
			}else{
				break
			}
		}
	}
	return -1
}
