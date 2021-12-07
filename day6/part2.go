package day6

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 6, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day6, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
	days := 256
	in := strings.Split(input.GetStringSlice()[0], ",")
	fishes := make([]int, 9)

	for _, fish := range in {
		fishes[util.AtoIorEXIT(fish)]++
	}
	for day := 0; day < days; day++ {
		newFishes := fishes[0]
		fishes = util.RotateLeft(fishes, 1)
		fishes[8] = newFishes
		fishes[6] += newFishes
	}

	sum := 0
	for _, value := range fishes {
		sum += value
	}
	return sum
}
