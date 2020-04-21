package leetcode

/*
	给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
	注意:
	每个数组中的元素不会超过 100
	数组的大小不会超过 200
	示例 1:
	输入: [1, 5, 11, 5]
	输出: true
	解释: 数组可以分割成 [1, 5, 5] 和 [11].
	 
	示例 2:
	输入: [1, 2, 3, 5]
	输出: false
	解释: 数组不能分割成两个元素和相等的子集.
*/

func canPartition(nums []int) bool {
	sum := 0
	out := false
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	partition(0, nums, target, &out)
	return out
}

func partition(index int, nums []int, target int, out *bool) {
	if *out == true {
		return
	}
	if target < 0 {
		return
	}
	if target == 0 {
		*out = true
		return
	}
	for i := index; i < len(nums); i++ {
		partition(i+1, nums, target-nums[i], out)
	}
}
