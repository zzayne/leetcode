package random_pick_with_blacklist

import "math/rand"

type Solution struct {
	bw    map[int]int
	bound int
}

func Constructor(n int, blacklist []int) Solution {
	m := len(blacklist)
	bound := n - m
	blackM := make(map[int]bool)
	for _, v := range blacklist {
		if v >= bound {
			blackM[v] = true
		}
	}
	bw := make(map[int]int, m-len(blackM))
	w := bound
	for _, v := range blacklist {
		if v < bound {
			for blackM[w] {
				w++
			}
			bw[v] = w
			w++
		}
	}
	return Solution{bw: bw, bound: bound}
}

func (s *Solution) Pick() int {
	x := rand.Intn(s.bound)
	if s.bw[x] > 0 {
		return s.bw[x]
	}
	return x
}
