package main

import (
	"fmt"
	testcase "komerce/test_case"
)

func main() {
	text := "Sample Case"
	vowels, consonants := testcase.VowelsAndConsonants(text)
	fmt.Println("Vowels:", vowels)
	fmt.Println("Consonants:", consonants)

	psbb := testcase.PSBB(5, []int{1, 2, 4, 3, 3})
	fmt.Println("PSBB:", psbb)

	psbb = testcase.PSBB(8, []int{2, 3, 4, 4, 2, 1, 3, 1})
	fmt.Println("PSBB:", psbb)
}
