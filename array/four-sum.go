package array

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}

	var res [][]int
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lastMatched := false
		for j := i + 1; j < n-2; j++ {
			left, right := j+1, n-1
			if lastMatched && nums[j] == nums[j-1] {
				continue
			}
			for left < right {
				n1, n2, n3, n4 := nums[i], nums[j], nums[left], nums[right]
				if n1+n2+n3+n4 == target {
					res = append(res, []int{n1, n2, n3, n4})
					lastMatched = true
					for left < right && nums[left] == n3 {
						left++
					}
					for left < right && nums[right] == n4 {
						right--
					}
				} else if n1+n2+n3+n4 < target {
					left++
				} else {
					right--
				}
			}

		}
	}
	return res
}
