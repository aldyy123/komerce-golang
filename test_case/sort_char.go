package testcase

import (
	"sort"
	"strings"
)

func sortSlice(s []rune) []rune {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return s
}

func VowelsAndConsonants(text string) (string, string) {
	var vowels string
	var consonants string

	for _, char := range strings.TrimSpace(strings.ToLower(text)) {
		if strings.Contains("aiueo", string(char)) {
			vowels += string(char)
		} else {
			if string(char) != " " {
				consonants += string(char)
			}
		}
	}

	return string(sortSlice([]rune(vowels))), string(sortSlice([]rune(consonants)))
}
