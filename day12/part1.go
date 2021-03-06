package day12

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 12, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day12, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	nodes := map[string][]string{}

	for line := range input.GetLines() {
		parts := strings.Split(line, "-")

		nodes[parts[0]] = append(nodes[parts[0]], parts[1])
		nodes[parts[1]] = append(nodes[parts[1]], parts[0])
	}

	return 0
}
