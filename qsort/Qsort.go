package main

import (
	"fmt"
)

//var count int = 0

func main() {
	input := []int{6, 5, 2, 7, 3, 9, 8, 4, 10, 1}
	fmt.Println("input:")
	fmt.Println(input)
	lenth := len(input) - 1

	qsort(input, 0, lenth)

	fmt.Println("after sort")
	fmt.Println(input)
}

func qsort(arr []int, start int, end int) {
	if start > end {
		return
	} else {
		index := partition(arr, start, end)
		qsort(arr, start, index-1)
		qsort(arr, index+1, end)
	}
}

func partition(arr []int, start int, end int) int {
	left := start
	right := end
	temp := arr[start]
	for left < right {
		for left < right && arr[right] >= temp {
			right -= 1
			//从右边找比temp小的
		}
		for left < right && arr[left] <= temp {
			left += 1
			//从左边找比temp大的
		}
		arr[right], arr[left] = arr[left], arr[right]
	}
	arr[start], arr[left] = arr[left], arr[start]
	//count++
	//fmt.Println("count: ", count)
	//fmt.Println("arr: ", arr)
	//fmt.Println("right:", right, ", left:", left)
	//return left

	return right
}

