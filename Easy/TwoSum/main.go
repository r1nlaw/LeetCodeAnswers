package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var res []int

	for index, value := range nums {
		mapValue, ok := m[target-value]
		if ok {
			res = append(res, index)
			res = append(res, mapValue)
		} else {
			m[value] = index
		}

	}

	return res

}