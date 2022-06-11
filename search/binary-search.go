package search

func search(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := (high-low)/2 + low
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
