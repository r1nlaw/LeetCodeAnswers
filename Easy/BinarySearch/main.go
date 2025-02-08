package main

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for i := 0; i < len(nums); i++ {
		mid := (low + high) / 2

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
		if nums[mid] == target {
			return mid

		}
	}

	return -1

}
