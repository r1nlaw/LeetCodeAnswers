package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		i := (low + high) / 2
		j := (m+n+1)/2 - i

		leftMax1 := math.Inf(-1)
		if i > 0 {
			leftMax1 = float64(nums1[i-1])
		}

		rightMin1 := math.Inf(1)
		if i < m {
			rightMin1 = float64(nums1[i])
		}

		leftMax2 := math.Inf(-1)
		if j > 0 {
			leftMax2 = float64(nums2[j-1])
		}

		rightMin2 := math.Inf(1)
		if j < n {
			rightMin2 = float64(nums2[j])
		}

		if leftMax1 <= rightMin2 && leftMax2 <= rightMin1 {
			if (m+n)%2 == 0 {
				return (math.Max(leftMax1, leftMax2) + math.Min(rightMin1, rightMin2)) / 2.0
			}
			return math.Max(leftMax1, leftMax2)
		} else if leftMax1 > rightMin2 {
			high = i - 1
		} else {
			low = i + 1
		}
	}

	return 0.0
}
