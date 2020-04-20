package leetcode

/*
	1. 两数之和
	给定一个整数数组 strs 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
	示例:
	给定 strs = [2, 7, 11, 15], target = 9
	因为 strs[0] + strs[1] = 2 + 7 = 9
	所以返回 [0, 1]
 */

func twoSum(nums []int, target int) []int {
	var targetMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := targetMap[nums[i]]; !ok {
			targetMap[target-nums[i]] = i
		}else {
			return []int{targetMap[nums[i]], i}
		}
	}
	return []int{}
}
