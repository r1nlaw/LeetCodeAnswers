package main

func arrangeCoins(n int) int {
	low, high := 1, n
	result := 0

	for low <= high {
		countRows := (low + high) / 2
		coins := (float64(countRows) / 2) * (float64(countRows) + 1)

		if coins > float64(n) {
			high = countRows - 1
		} else {
			low = countRows + 1
			result = max(countRows, result)
		}
	}
	return result

}
