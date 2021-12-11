package main

func minWindow(s string, t string) string {
	mt, ms := map[byte]int{}, map[byte]int{}

	var matchLen int
	var ans string

	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}

	for l, r := 0, 0; r < len(s); r++ {
		ms[s[r]]++
		if _, ok := mt[s[r]]; ok && ms[s[r]] == mt[s[r]] {
			matchLen++
		}
		for matchLen == len(mt) {
			if len(ans) == 0 || len(ans) > len(s[l:r+1]) {
				ans = s[l : r+1]
			}
			if _, ok := mt[s[l]]; ok && ms[s[l]] == mt[s[l]] {
				matchLen--
			}
			ms[s[l]]--
			l++
		}
	}
	return ans
}
