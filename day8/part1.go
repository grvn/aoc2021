package day8

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

const (
	// antal segment
	ONE   = 2
	FOUR  = 4
	SEVEN = 3
	EIGHT = 7
)

var NUMBERS = map[int]bool{
	ONE:   true,
	FOUR:  true,
	SEVEN: true,
	EIGHT: true,
}

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 8, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day8, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	sum := 0
	for line := range input.GetLines() {
		splitLine := strings.Split(line, " | ")
		rightSideParts := strings.Fields(splitLine[1])
		for _, part := range rightSideParts {
			if NUMBERS[len(part)] {
				sum++
			}
		}
	}
	return sum
}
