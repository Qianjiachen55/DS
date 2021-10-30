package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	if totalLength%2 == 1 {
		// 5-->2
		midIndex := totalLength / 2

		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		//4->1,2
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex)+getKthElement(nums1, nums2, midIndex+1)) / 2.0
	}

	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1

		p1, p2 := nums1[newIndex1], nums2[newIndex2]
		if p1 <= p2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}

	}
	return 0

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	res := findMedianSortedArrays(nums1, nums2)

	fmt.Println(res)
}