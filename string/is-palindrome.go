package string

// https://leetcode.cn/problems/valid-palindrome

// time cost: beyond 100.00%
// mem cost:  beyond 83.87%
func isPalindrome(s string) bool {
	isValid := func(s uint8) bool {
		if s >= 97 && s <= 122 {
			return true
		}
		if s >= 65 && s <= 90 {
			return true
		}
		if s >= 48 && s <= 57 {
			return true
		}
		return false
	}

	lower := func(s uint8) uint8 {
		if s >= 65 && s <= 90 {
			s += 32
		}

		return s
	}

	i := 0
	j := len(s) - 1

	for i < j {
		if !isValid(s[i]) {
			i++
			continue
		}
		if !isValid(s[j]) {
			j--
			continue
		}
		if lower(s[i]) != lower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

// time cost: beyond 100.00%
// mem cost:  beyond 18.82%
func isPalindrome2(s string) bool {
	var lst []rune
	for _, v := range []rune(s) {
		if (v >= '0' && v <= '9') || (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			if v >= 'A' && v <= 'Z' {
				lst = append(lst, v+32)
			} else {
				lst = append(lst, v)
			}
		}
	}

	if len(lst) < 2 {
		return true
	}
	length := len(lst)
	for i := 0; i < length/2; i++ {
		if lst[i] != lst[length-i-1] {
			return false
		}
	}
	return true
}
