package main


func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)==0{
		if s2 ==s3{
			return true
		}
		return false
	}
	if len(s2) == 0{
		if s3==s1{
			return true
		}
		return false
	}
	if len(s1)+len(s2) !=len(s3){
		return false
	}

	ans :=make([][]bool,len(s2)+1)
	for i:=0;i<len(s2)+1;i++{
		ans[i] = make([]bool,len(s1)+1)
		if s2[:i]==s3[:i]{
			ans[i][0]=true
		}
	}
	for j:=0;j<len(s1)+1;j++{
		if s1[:j]==s3[:j]{
			ans[0][j]=true
		}
	}

	for i:=1;i<=len(s2);i++{
		for j:=1;j<=len(s1);j++{
			if ans[i-1][j] && s2[i-1]==s3[i+j-1]{
				ans[i][j]=true
			}
			if ans[i][j-1] && s1[j-1]==s3[i+j-1]{
				ans[i][j]=true
			}
		}
	}

	return ans[len(s2)][len(s1)]


}
