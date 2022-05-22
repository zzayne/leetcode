package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var sameIdx []int
	for i, v := range nums {
		if v == target-v {
			sameIdx = append(sameIdx, i)
			continue
		}
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}

	if len(sameIdx) > 0 {
		return sameIdx
	}

	for num, idx := range m {
		if j, ok := m[target-num]; ok {
			return []int{idx, j}
		}
	}
	return nil
}
