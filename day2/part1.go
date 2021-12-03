package day2

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 2, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day2, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	instructions := input.GetLines()
	x, depth := 0, 0

	for instruction := range instructions {
		instructionparts := strings.Split(instruction, " ")
		direction, value := instructionparts[0], util.AtoIorEXIT(instructionparts[1])
		switch direction {
		case UP:
			depth -= value
		case DOWN:
			depth += value
		case FORWARD:
			x += value
		}
	}
	result := depth * x
	return result
}
