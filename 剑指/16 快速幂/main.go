package main


func myPow(x float64, n int) float64 {
	isZ := true
	if n <0{
		isZ = false
		n = -n
	}
	res := 1.0
	for n >0{
		if n %2 ==1{
			res = res * x
		}
		x = x*x
		n = n/2
	}
	if !isZ{
		return 1.0/res
	}
	return res
}
