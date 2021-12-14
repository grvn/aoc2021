package day11

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 11, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day11, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
	return 0
}
