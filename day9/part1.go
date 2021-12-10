package day9

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 9, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day9, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	return 0
}
