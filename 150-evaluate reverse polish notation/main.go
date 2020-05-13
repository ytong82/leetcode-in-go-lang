package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			a, b := nums[len(nums)-2], nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, compute(a, b, token))
		default:
			temp, _ := strconv.Atoi(token)
			nums = append(nums, temp)
		}
	}
	return nums[0]
}

func compute(a, b int, opt string) int {
	switch opt {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))

	tokens = []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))

	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
