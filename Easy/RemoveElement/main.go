package main

import "fmt"

func removeByElement(array []int, index int) ([]int, error) {
	if index >= len(array) || index < 0 {
		return nil, fmt.Errorf("Index is out if range", index, len(array))
	}

	array[index] = array[len(array)-1]

	return array[:len(array)-1], nil
}

func removeElement(nums []int, val int) int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == val {
			var err error
			nums, err = removeByElement(nums, i)
			if err != nil {
				fmt.Println("Error", err)
				return len(nums)
			}
			i--
		}

	}
	fmt.Println(nums)
	return len(nums)

}

func main() {
	var nums = []int{3, 2, 2, 3}
	var val = 3

	removeElement(nums, val)
}
