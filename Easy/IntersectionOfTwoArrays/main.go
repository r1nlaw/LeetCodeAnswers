package main

func minArr(nums1 []int, nums2 []int) ([]int, []int) {
	if len(nums1) < len(nums2) {
		return nums1, nums2
	} else {
		return nums2, nums1
	}
}

func intersection(nums1 []int, nums2 []int) []int {

	probablyResult := []int{}

	minArr, maxArr := minArr(nums1, nums2)

	for i := 0; i < len(minArr); i++ {
		for j := 0; j < len(maxArr); j++ {
			if minArr[i] == maxArr[j] {
				probablyResult = append(probablyResult, minArr[i])
			}
		}
	}

	uniqueMap := make(map[int]bool)
	var result []int

	for _, val := range probablyResult {
		if !uniqueMap[val] {
			uniqueMap[val] = true
			result = append(result, val)
		}
	}

	return result
}
