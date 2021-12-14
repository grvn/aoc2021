package util

import (
	"os"
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func SortString(input string) string {
	s := []rune(input)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return string(s)
}

// string becomes int or everything crashes
func AtoIorEXIT(s string) int {
	output, err := strconv.Atoi(s)
	if err != nil {
		log.WithFields((log.Fields{
			"Error": err,
		})).Error("line can not be converted to int")
		os.Exit(1)
	}
	return output
}
