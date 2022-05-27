package string

//https://leetcode.cn/problems/valid-palindrome/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sLst, tLst := make([]int, 26), make([]int, 26)
	for _, v := range s {
		sLst[v-'a']++
	}
	for _, v := range t {
		tLst[v-'a']++
	}
	for i := 0; i < len(sLst); i++ {
		if sLst[i] != tLst[i] {
			return false
		}
	}
	return true
}
