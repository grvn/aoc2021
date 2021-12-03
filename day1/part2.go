package day1

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 1, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day1, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
	depths := input.GetIntSlice()
	result, prev := 0, 0
	for i := 0; i+2 < len(depths); i++ {
		curr := depths[i] + depths[i+1] + depths[i+2]
		if curr > prev && prev != 0 {
			result++
		}
		prev = curr
	}
	return result
}
