package main

func isPerfectSquare(num int) bool {

	if num < 0 {
		return false
	}
	if num == 0 || num == 1 {
		return true
	}

	x := float64(num)
	guess := x / 2

	for i := 0; i < 20; i++ {
		guess = (guess + x/guess) / 2
	}

	sqrt := int(guess)
	return sqrt*sqrt == num

}
