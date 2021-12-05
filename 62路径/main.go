package main




//func uniquePaths(m int, n int) int {
//
//	return cFun(n+m-2,m-1)
//}
//
//func cFun(a,b int)int{
//	x :=1
//	y:=1
//	count := min(b,a-b)
//	for i:=0;i<count;i++{
//		x *=a
//		a--
//	}
//	for i:=1;i<=count;i++{
//		y*=i
//	}
//	return x/y
//
//}
//
//func min(a,b int)int{
//	if a>b{
//		return b
//	}
//	return a
//}







/////////////////////////

//2
func uniquePaths(m int, n int) int {

	return cFun(n+m-2,m-1)
}

func cFun(a,b int)int{
	x :=1
	y:=1
	count := min(b,a-b)
	for i:=0;i<count;i++{
		x *=a
		a--
	}
	for i:=1;i<=count;i++{
		y*=i
	}
	return x/y

}

func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}