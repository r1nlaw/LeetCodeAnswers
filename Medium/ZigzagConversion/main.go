package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := ""

	for r := 0; r < numRows; r++ {
		increment := 2 * (numRows - 1)
		for i := r; i < len(s); i += increment {
			result += string(s[i])
			if r > 0 && r < numRows-1 && i+increment-2*r < len(s) {
				result += string(s[i+increment-2*r])
			}
		}
	}
	return result
}
