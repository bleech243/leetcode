package leetcode

/*
	地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，
	它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
	例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，
	因为3+5+3+8=19。请问该机器人能够到达多少个格子？

	示例 1：
	输入：m = 2, n = 3, k = 1
	输出：3
	示例 2：
	输入：m = 3, n = 1, k = 0
	输出：1
	提示：
	1 <= n,m <= 100
	0 <= k <= 20
*/

func movingCount(m int, n int, k int) int {
	grids := make([][]int, m+2)
	for i := 0; i < m+2; i++ {
		grids[i] = make([]int, n+2)
		for j := 0; j < n+2; j++ {
			if i == 0 || i == m+1 || j == 0 || j == n+1 || isBarrier(i-1, j-1, k) {
				grids[i][j] = 2
			}else{
				grids[i][j] = 0
			}
		}
	}

	move(1, 1, grids)
	count := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if grids[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func move(i int, j int, grids [][]int) {
	if grids[i][j] == 1 {
		return
	}
	if grids[i][j] == 0 {
		grids[i][j] = 1
	}
	if grids[i+1][j] != 2 {
		move(i+1, j, grids)
	}
	if grids[i][j+1] != 2 {
		move(i, j+1, grids)
	}
	if grids[i-1][j] != 2 {
		move(i-1, j, grids)
	}
	if grids[i][j-1] != 2 {
		move(i, j-1, grids)
	}
}

func isBarrier(m int, n int, k int) bool {
	return getTotal(m)+getTotal(n) > k
}

func getTotal(i int) int {
	return i/100%10 + i/10%10 + i%10
}
