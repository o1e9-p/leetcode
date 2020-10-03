package spiralMatrix

// https://leetcode.com/problems/spiral-matrix/
func SpiralOrder(matrix [][]int) []int {
	var(
		m = len(matrix)
		n = len(matrix[0])
		output = make([]int, 0, m*n)
		elAmount = m*n
	)

	for i,j := 0,0; i < m && j < n; {
		for a := j; a < n && elAmount > len(output); a++ {
			output = append(output, matrix[i][a])
		}
		i++

		for b := i; b < m && elAmount > len(output); b++ {
			output = append(output, matrix[b][n-1])
		}
		n--

		for c := n-1; c >= j && elAmount > len(output); c-- {
			output = append(output, matrix[m-1][c])
		}
		m--

		for d := m-1; d >= i && elAmount > len(output); d-- {
			output = append(output, matrix[d][j])
		}
		j++
	}
	return output
}
