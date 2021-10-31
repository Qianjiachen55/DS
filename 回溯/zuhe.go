package main

var hash = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var ans []string

func letterCombinations(digits string) []string {

	ans = []string{}

	if len(digits) == 0 {
		return []string{}
	}

	backTrack(digits, 0, "")
	return ans

}

func backTrack(digits string, index int, s string) {
	if index == len(digits) {
		ans = append(ans, s)
	} else {
		digit := string(digits[index])
		zimu := hash[digit]
		zimuCount := len(zimu)
		for i := 0; i < zimuCount; i++ {
			backTrack(digits, index+1, s+string(zimu[i]))
		}

	}
}

//涉及递归传参数，过度全局变量
//注意类型 char取出来可能int 了
//建议string存储
