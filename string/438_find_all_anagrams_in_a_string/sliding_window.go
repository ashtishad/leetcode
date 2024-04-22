package main

import (
	"fmt"
	"slices"
)

func findAnagrams(s string, p string) (res []int) {
	countP := make([]rune, 26)
	countS := make([]rune, 26)

	for i := range p {
		countP[p[i]-'a']++
	}

	l := 0
	for r := range s {
		countS[s[r]-'a']++

		if r-l+1 == len(p) {
			if slices.Equal(countS, countP) {
				res = append(res, l)
			}

			countS[s[l]-'a']--
			l++
		}
	}

	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
