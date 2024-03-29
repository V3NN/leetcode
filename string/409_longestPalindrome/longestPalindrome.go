package __longestPalindrome

/*
  409. 最长回文串

  给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串
  在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。
*/

func longestPalindrome(s string) int {
	res := 0
	hash := map[rune]int{}
	for _, v := range s {
		hash[v]++
	}

	for _, v := range hash {
		if v%2 == 0 {
			res += v
		} else {
			res += v - 1
		}
	}

	if len(s) == res {
		return res
	}

	return res + 1
}