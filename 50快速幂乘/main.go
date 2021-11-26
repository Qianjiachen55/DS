package main

func myPow(x float64, n int) float64 {
	if n>=0{
		return Ppow(x,n)
	}else{
		return 1.0/Ppow(x,-n)
	}
}

func Ppow(x float64,n int)float64{
	if n==1{
		return x
	}
	if n==0{
		return 1
	}

	cur := Ppow(x,n/2)
	if n%2 ==1{
		return cur * cur * x
	}
	return cur *cur
}
