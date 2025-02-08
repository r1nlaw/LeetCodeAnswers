package main

func mySqrt(x int) int {
	low := 0
	high := x

	for low <= high {
		mid := (low + high) / 2

		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high

}
