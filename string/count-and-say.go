package string

import (
	"strconv"
)

//https://leetcode.cn/problems/count-and-say/

func countAndSay(n int) string {
	current := "1"
	if n == 1 {
		return current
	}
	do := func(str string) string {
		res := ""
		for i := 0; i < len(str); i++ {
			temp := 1
			num := str[i : i+1]
			for k := i + 1; k < len(str); k++ {
				if str[i] == str[k] {
					temp++
				}
				if str[i] != str[k] {
					i = k - 1
					break
				}
				if k == len(str)-1 {
					i = k
				}
			}
			res = res + strconv.Itoa(temp) + num
		}
		return res
	}

	for i := 1; i < n; i++ {
		current = do(current)
	}
	return current
}
