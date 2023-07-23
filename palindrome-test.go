package main

import (
	"fmt"
	"strings"
)

func isPalindrome(data string) bool {
	data = strings.ToLower(strings.ReplaceAll(data, " ", ""))
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		if data[i] != data[j] {
			return false
		}
	}
	return true
}

func main() {
	input := "SAIPPUAKIVIKAUPPIAS"
	if isPalindrome(input) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
