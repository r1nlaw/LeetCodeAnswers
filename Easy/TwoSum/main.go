package main

func twoSum(nums []int, target int) []int {

	var array []int

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+nums[i+1] == target {
			array = append(array, i)
			array = append(array, i+1)
		}
	}

	return array
}
