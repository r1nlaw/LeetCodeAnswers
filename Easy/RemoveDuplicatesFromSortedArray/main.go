package main

import "fmt"

func removeDuplicates(nums []int) int {

	counter := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[counter] = nums[i]
			counter++
		}
	}
	fmt.Print(counter)
	return counter

}
