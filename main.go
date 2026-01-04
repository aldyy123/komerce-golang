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

	vowels, consonants = testcase.VowelsAndConsonants("Next Case")
	fmt.Println("Vowels:", vowels)
	fmt.Println("Consonants:", consonants)

	fmt.Println("--------------------------------")

	psbb := testcase.PSBB(5, []int{1, 2, 4, 3, 3})
	fmt.Println("PSBB 1:", psbb)

	psbb = testcase.PSBB(8, []int{2, 3, 4, 4, 2, 1, 3, 1})
	fmt.Println("PSBB 2:", psbb)

	psbb = testcase.PSBB(5, []int{15})
	fmt.Println("PSBB 3:", psbb)
}
