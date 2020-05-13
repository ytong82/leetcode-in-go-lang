package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	dict := make(map[string]bool)
	res := make([]string, 0)

	for i := 0; i <= len(s)-10; i++ {
		begin := i
		end := begin + 10
		key := s[begin:end]

		in, ok := dict[key]
		if ok {
			if in {
				res = append(res, key)
			}
			dict[key] = false
		} else {
			dict[key] = true
		}
	}
	return res
}

func main() {
	ss := []string{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", "AAAAAAAAAAAA"}
	for _, s := range ss {
		seqs := findRepeatedDnaSequences(s)
		for _, seq := range seqs {
			fmt.Println(seq)
		}
	}
}
