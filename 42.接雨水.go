package leetcode

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）
*/

func trap(height []int) int {
	result := 0
	for i, j := 0, len(height)-1; i < j; {
		if i == j {
			return result
		}
		if height[i] < height[j] {
			m := i + 1
			for ; height[m] < height[i]; m++ {
				result += height[i] - height[m]
			}
			i = m
		} else {
			n := j - 1
			for ; height[n] < height[j]; n-- {
				result += height[j] - height[n]
			}
			j = n
		}
	}
	return result
}
