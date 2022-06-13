package string

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}
	if s1 == s2 {
		return true
	}

	var chars1, chars2 [26]int
	for i, ch := range s1 {
		chars1[ch-'a']++
		chars2[s2[i]-'a']++
	}
	if chars1 == chars2 {
		return true
	}

	for i := n1; i < n2; i++ {
		chars2[s2[i]-'a']++
		chars2[s2[i-n1]-'a']--
		if chars1 == chars2 {
			return true
		}
	}

	return false
}
