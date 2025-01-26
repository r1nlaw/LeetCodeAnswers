package main

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}
	result := ""
	for i := 0; i < minLength; i++ {
		for _, str := range strs {
			if str[i] != strs[0][i] {
				return result
			}
		}
		result += string(strs[0][i])
	}

	return result
}
