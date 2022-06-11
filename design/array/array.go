package array

import "math/rand"

type Solution struct {
	nums []int
}

func SolutionConstructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}

}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	res := make([]int, len(this.nums))
	copy(res, this.nums)
	for i := 0; i < len(res); i++ {
		j := rand.Intn(i + 1)
		res[i], res[j] = res[j], res[i]
	}
	return res
}
