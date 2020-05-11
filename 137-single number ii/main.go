package main

import "fmt"

func singleNumber(nums []int) int {
	dict := make(map[int]int)
	sum := 0
	setsum := 0
	for _, num := range nums {
			_, ok := dict[num]
			if !ok {
				setsum = setsum + num
				dict[num] = 1
			}
			sum = sum + num
	}
	return (setsum * 3 - sum) / 2
}

func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))

	nums = []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}