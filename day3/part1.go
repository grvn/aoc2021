package day3

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 3, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day3, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	// https://yourbasic.org/golang/operators/
	lines := input.GetIntMatrix()
	tlines := util.Transpose(lines)
	epsilon, gamma := 0, 0
	for _, line := range tlines {
		// left shift ( add 0 to both)
		epsilon <<= 1
		gamma <<= 1
		switch most(line) {
		case 0:
			// bitwise OR
			epsilon |= 1
		case 1:
			gamma |= 1
		}
	}
	return epsilon * gamma
}

func most(input []int) int {
	zeros, ones := 0, 0
	for _, value := range input {
		switch value {
		case 0:
			zeros++
		case 1:
			ones++
		}
	}
	if zeros > ones {
		return 0
	} else {
		return 1
	}
}
