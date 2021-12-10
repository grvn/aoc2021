package util

import "sort"

func SortString(input string) string {
	s := []rune(input)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return string(s)
}
