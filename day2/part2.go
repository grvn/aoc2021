package day2

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 2, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day2, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
	instructions := input.GetLines()
	aim, x, depth := 0, 0, 0

	for instruction := range instructions {
		instructionparts := strings.Split(instruction, " ")
		direction, value := instructionparts[0], util.AtoIorEXIT(instructionparts[1])
		switch direction {
		case UP:
			aim -= value
		case DOWN:
			aim += value
		case FORWARD:
			x += value
			depth += aim * value
		}
	}
	result := depth * x
	return result
}
