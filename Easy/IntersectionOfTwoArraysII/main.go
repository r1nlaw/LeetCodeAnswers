package main

func intersect(nums1 []int, nums2 []int) []int {
	counts := make(map[int]int)
	var output []int

	for _, num := range nums1 {
		counts[num]++
	}

	for _, num := range nums2 {
		if counts[num] > 0 {
			output = append(output, num)
			counts[num]--
		}
	}

	return output
}
