package day11

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

type point struct {
	x, y int
}

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 11, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day11, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	flashes := 0
	lines := input.GetStringSlice()
	octopuses := make(map[point]int)
	for x, line := range lines {
		for y, value := range line {
			octopuses[point{x: x, y: y}] = int(value - '0')
		}
	}
	for i := 0; ; i++ {
		// Step 1
		for octopus := range octopuses {
			octopuses[octopus]++
		}

		// Step 2
		flashqueue := util.NewQueue()
		flashed := make(map[point]bool)
		for octopus := range octopuses {
			if octopuses[octopus] > 9 {
				flashes++
				flashed[octopus] = true
				for _, neighoctopus := range getNeighbours(octopuses, octopus) {
					octopuses[neighoctopus]++
					if octopuses[neighoctopus] > 9 {
						if _, ok := flashed[neighoctopus]; !ok {
							flashqueue.Push(neighoctopus)
						}
					}
				}
			}
		}
		for flashqueue.Length() > 0 {
			octo, _ := flashqueue.Pop()
			octopus := octo.(point)
			_, ok := flashed[octopus]
			if octopuses[octopus] > 9 && !ok {
				flashes++
				flashed[octopus] = true
				for _, octo := range getNeighbours(octopuses, octopus) {
					octopuses[octo]++
					if octopuses[octo] > 9 {
						if _, ok := flashed[octo]; !ok {
							flashqueue.Push(octo)
						}
					}
				}
			}
		}

		// Step 3
		for octopus := range flashed {
			octopuses[octopus] = 0
		}

		if i == 99 {
			break
		}
	}

	return flashes
}
