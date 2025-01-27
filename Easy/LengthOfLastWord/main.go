package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {

	s = strings.TrimSpace(s)
	words := strings.Fields(s)

	if len(words) > 0 {
		return len(words[len(words)-1])
	}

	return 0

}
