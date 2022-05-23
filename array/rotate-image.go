package array

// https://leetcode.cn/problems/rotate-image/
func rotate(matrix [][]int) {
	length := len(matrix)

	for i := 0; i < length/2; i++ {
		for j := i; j < length-i-1; j++ {
			m := length - j - 1
			n := length - i - 1
			tmp := matrix[i][j]
			matrix[i][j] = matrix[m][i]
			matrix[m][i] = matrix[n][m]
			matrix[n][m] = matrix[j][n]
			matrix[j][n] = tmp
		}

	}
}
