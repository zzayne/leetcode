package array

// https://leetcode.cn/problems/intersection-of-two-arrays-ii/
func intersect(nums1 []int, nums2 []int) []int {
	var out []int

	m := map[int]int{}
	for _, i := range nums1 {
		v := m[i]
		v++
		m[i] = v
	}
	for _, i := range nums2 {
		v := m[i]
		if v > 0 {
			out = append(out, i)
			v--
			m[i] = v
		}
	}
	return out
}
