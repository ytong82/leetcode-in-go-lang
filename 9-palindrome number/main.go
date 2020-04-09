package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	length := len(s)

	mid := length / 2
	var left, right int
	if length % 2 == 0 {
		right = mid 
		left = mid  - 1
	} else {
		right = mid + 1
		left = mid - 1
	}

	for left >= 0 && right < length {
		if s[left] != s[right] {
			return false
		}
		left = left - 1
		right = right + 1
	}

	return true
}

func main() {
	input := 121
	res := isPalindrome(input)
	fmt.Printf("%v is a palindrome? %v \n", input, res)

	input = -121
	res = isPalindrome(input)
	fmt.Printf("%v is a palindrome? %v \n", input, res)

	input = 10
	res = isPalindrome(input)
	fmt.Printf("%v is a palindrome? %v \n", input, res)
} 