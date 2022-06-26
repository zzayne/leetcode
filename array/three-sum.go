package array

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	var res [][]int

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right {
			n1, n2, n3 := nums[i], nums[left], nums[right]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for left < right && nums[left] == n2 {
					left++
				}
				for left < right && nums[right] == n3 {
					right--
				}
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
