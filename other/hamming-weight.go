package other

func hammingWeight(num uint32) int {
	i := 0
	for num != 0 {
		if num%2 == 1 {
			i++
		}
		num = num >> 1
	}
	return i
}

unc minimumTotal(triangle [][]int) int {
if len(triangle) < 1 {
return 0
}
//从倒数第 2 层推
for i := len(triangle) - 2; i >= 0; i-- {
for j := 0; j < len(triangle[i]); j++ {
triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
}
}
return triangle[0][0]
}
