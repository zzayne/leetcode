package string

// brute force
func strStr(haystack string, needle string) int {
	len1 := len(haystack)
	len2 := len(needle)
	if haystack == needle {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if len1-i < len2 {
			return -1
		}
		if haystack[i:i+len2] == needle {
			return i
		}
	}
	return -1
}
