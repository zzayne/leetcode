package search

func firstBadVersion(n int) int {
	res, low, high := 0, 0, n
	for low <= high {
		mid := (high-low)/2 + low
		if !isBadVersion(mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
		res = low
	}
	return res
}

func isBadVersion(i int) bool {
	return false
}
