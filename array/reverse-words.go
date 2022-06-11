package array

func reverseWords(s string) string {
	n := len(s)
	var ret []byte
	for i := 0; i < n; {
		start := i
		for i < n && s[i] != ' ' {
			i++
		}
		for p := i - 1; p >= start; p-- {
			ret = append(ret, s[p])
		}
		for i < n && s[i] == ' ' {
			ret = append(ret, ' ')
			i++
		}
	}
	return string(ret)
}
