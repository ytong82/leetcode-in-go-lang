package main

import "fmt"

func singleNumber(nums []int) int {
	dict := make(map[int]int)
	for _, num := range nums {
			dict[num] = dict[num] + 1
			if dict[num] == 3 {
				delete(dict, num)
			}
	}

	res := 0
	for key, _ := range dict {
		res = key
	}
	return res
}

func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))

	nums = []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}