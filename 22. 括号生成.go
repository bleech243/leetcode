package leetcode

/*
	22. 括号生成
	数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

	示例：
	输入：n = 3
	输出：[
		   "((()))",
		   "(()())",
		   "(())()",
		   "()(())",
		   "()()()"
		 ]
*/

func generateParenthesis(n int) []string {
	var result []string
	countLeft := 1
	countRight := 0
	addParenthesis(countLeft, countRight, n, "(", &result)
	return result
}

func addParenthesis(countLeft int, countRight int, n int, str string, result *[]string) {
	if countLeft == n && countLeft == countRight {
		*result = append(*result, str)
	}
	if countLeft <= n && countRight <= countLeft {
		if countRight < countLeft {
			addParenthesis(countLeft+1, countRight, n, str+"(", result)
			addParenthesis(countLeft, countRight+1, n, str+")", result)
		} else if countRight == countLeft {
			addParenthesis(countLeft+1, countRight, n, str+"(", result)
		}
	}
}
