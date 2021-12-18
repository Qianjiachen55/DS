package main


func merge(nums1 []int, m int, nums2 []int, n int)  {
	cur := m+n-1
	if m==0 && n==0{
		nums1 = []int{}
		return
	}
	pm := m-1
	pn := n-1
	for pm>=0 || pn>=0{
		if pn<0{
			nums1[cur] = nums1[pm]
			pm--
			cur--
			continue
		}
		if pm <0{
			nums1[cur]=nums2[pn]
			pn--
			cur--
			continue
		}

		if nums1[pm]>nums2[pn]{
			nums1[cur] = nums1[pm]
			pm--
		}else{
			//nums1[pm]<=nums2[pn] || pm<0
			nums1[cur] = nums2[pn]
			pn--
		}
		cur--
	}
}
