package utils

/* 最长公共子串 */
func LCS(str1 string, str2 string) string {
	// write code here
	l1 := len(str1) //行
	l2 := len(str2) // 列
	if l1 == 0 || l2 == 0 {
		return ""
	}

	// 第一行 第一列都为空
	//   主要是为了好处理 dp[0][0]
	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}

	max := 0
	end := 0

	//  注意 l1+1   l2+1
	//  二维数组dp[i][j]表示第一个字符串前i个字符和第二个字符串前j个字符组成的最长公共字符串的长度
	//  字符相等 dp[i][j] = dp[i-1][j-1] + 1  否则 dp[i][j] = 0
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}

			if dp[i][j] > max {
				max = dp[i][j]
				end = i // 注意
			}
		}
	}

	if max == 0 {
		return ""
	}

	return str1[end-max : end]
}
