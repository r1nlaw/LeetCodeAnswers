package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	var res []string

	for i < len(word1) && j < len(word2) {
		res = append(res, string(word1[i]))
		res = append(res, string(word2[j]))
		i++
		j++

	}
	res = append(res, word1[i:])
	res = append(res, word2[j:])
	str := strings.Join(res, "")
	return str

}


