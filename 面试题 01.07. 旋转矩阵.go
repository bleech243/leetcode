package leetcode

func rotate(matrix [][]int) {

	var tmp1 int
	var tmp2 int
	var tmp3 int
	N := len(matrix)
	if N == 1 || N == 0 {
		return
	}
	for i := 0; i < N-1; i++ {
		tmp1 = matrix[i][0]
		tmp2 = matrix[0][N-i-1]
		tmp3 = matrix[N-i-1][N-1]
		matrix[i][0] = matrix[N-1][i]
		matrix[0][N-i-1] = tmp1
		matrix[N-i-1][N-1] = tmp2
		matrix[N-1][i] = tmp3
	}
	newMatrix := make([][]int, N-2)
	for i := 0; i < N-2; i++ {
		newMatrix[i] = matrix[i+1][1:N-1]
	}
	rotate(newMatrix)
}