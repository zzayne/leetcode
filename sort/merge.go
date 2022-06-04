package sort

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
	return
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	var i, j, idx int
	lst := make([]int, m+n)

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			lst[idx] = nums1[i]
			i++
		} else {
			lst[idx] = nums2[j]
			j++
		}
		idx++
	}
	for i < m {
		lst[idx] = nums1[i]
		idx++
		i++
	}
	for j < n {
		lst[idx] = nums2[j]
		idx++
		j++
	}
	for k := 0; k < m+n; k++ {
		nums1[k] = lst[k]
	}
	return
}
