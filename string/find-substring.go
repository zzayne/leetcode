package string

func findSubstring(s string, words []string) (ans []int) {
	ls, lw, lc := len(s), len(words), len(words[0])

	for i := 0; i < lc && i+lw*lc <= ls; i++ {
		wordHash := map[string]int{}
		for j := 0; j < lw; j++ {
			wordHash[s[i+j*lc:i+(j+1)*lc]]++
		}

		for _, word := range words {
			wordHash[word]--
			if wordHash[word] == 0 {
				delete(wordHash, word)
			}
		}

		for start := i; start < ls-lw*lc+1; start += lc {
			if start != i {
				word := s[start+(lw-1)*lc : start+lw*lc]
				wordHash[word]++
				if wordHash[word] == 0 {
					delete(wordHash, word)
				}
				word = s[start-lc : start]
				wordHash[word]--
				if wordHash[word] == 0 {
					delete(wordHash, word)
				}
			}

			if len(wordHash) == 0 {
				ans = append(ans, start)
			}
		}

	}
	return ans

}
