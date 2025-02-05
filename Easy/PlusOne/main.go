package main

import "strconv"

func plusOne(digits []int) []int {

	num := 0

	for _, val := range digits {
		num = num*10 + val
	}
	num++

	strNum := strconv.Itoa(num)

	newDigits := make([]int, len(strNum))

	for i, char := range strNum {
		newDigits[i] = int(char - '0')
	}

	return newDigits

}
