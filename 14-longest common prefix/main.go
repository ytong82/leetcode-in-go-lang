package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	minlen := math.MaxInt64

	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	for _, str := range strs {
		if len(str) < minlen {
			minlen = len(str)
		}
	}

	for i:=0; i<minlen; i++ {
		for j:=0; j<len(strs)-1; j++ {
			if strs[j][i] != strs[j+1][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:minlen]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Printf("Longest common prefix of array %v is %v.\n", strs, res)

	strs = []string{"dog", "racecar", "car"}
	res = longestCommonPrefix(strs)
	fmt.Printf("Longest common prefix of array %v is %v.\n", strs, res)

	strs = []string{}
	res = longestCommonPrefix(strs)
	fmt.Printf("Longest common prefix of array %v is %v.\n", strs, res)

	strs = []string{"a"}
	res = longestCommonPrefix(strs)
	fmt.Printf("Longest common prefix of array %v is %v.\n", strs, res)

	strs = []string{"a", "b"}
	res = longestCommonPrefix(strs)
	fmt.Printf("Longest common prefix of array %v is %v.\n", strs, res)

	strs = []string{"abc", "abc"}
	res = longestCommonPrefix(strs)
	fmt.Printf("Longest common prefix of array %v is %v.\n", strs, res)
}