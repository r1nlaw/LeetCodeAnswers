package main

func reverse(number int) int {
	newNumber := 0
	for number > 0 { 				// 121  12	1
		remainder := number % 10 	// 1	2	1
		newNumber *= 10          	// 0	10	120
		newNumber += remainder   	// 1	12	121
		number /= 10             	// 12	1	0
	}
	return newNumber
}

func isPalindrome(x int) bool {

	if x == reverse(x) {
		return true
	} else {
		return false
	}

}
