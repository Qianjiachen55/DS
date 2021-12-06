package main



func addBinary(a string, b string) string {
	add := false
	ans :=""
	// temp :=""

	i,j := len(a)-1,len(b)-1
	for i>=0||j>=0{
		m ,n := '0','0'
		count:=0
		if i>=0{
			m =rune(a[i])
			i--
		}
		if j>=0{

			n = rune(b[j])
			j--
		}
		if m=='1'{
			count ++
		}
		if n=='1'{
			count ++
		}
		if add {
			count ++
		}
		switch count{
		case 0:
			ans = "0"+ans
			add = false

			break
		case 1:
			ans = "1"+ans
			add = false
			break
		case 2:
			ans = "0"+ans
			add = true
			break
		case 3:
			ans = "1" + ans
			add = true
			break
		}


	}
	if add{
		ans = "1"+ans
	}

	if ans[0]=='0'&&len(ans)>1{
		return ans[1:]
	}
	return ans

}
