package day1

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 1, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day1, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	depths := input.GetIntSlice()
	result := 0
	for i := 1; i < len(depths); i++ {
		if i >= 1 && depths[i] > depths[i-1] {
			result++
		}
	}
	return result
}
