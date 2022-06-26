package string

func replaceSpace(s string) string {
	var out string
	for i, v := range s {
		if v-' ' == 0 {
			out = out + "%20"
			continue
		}
		out = out + string(s[i])
	}
	return out
}
