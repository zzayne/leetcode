package other

func generateTriangle(numRow int) [][]int {
	if numRow < 1 {
		return nil
	}
	var res [][]int
	res = append(res, []int{1})
	if numRow < 2 {
		return res
	}

	res = append(res, []int{1, 1})
	for i := 2; i < numRow; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			left := res[i-1][j-1]
			right := res[i-1][j]
			temp[j] = left + right
		}
		res = append(res, temp)
	}
	return res

}
