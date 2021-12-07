package day7

import (
	"math"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 7, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day7, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	min, max := math.MaxInt32, math.MinInt32
	in := strings.Split(input.GetStringSlice()[0], ",")
	crabs := make([]int, len(in))
	for i, crab := range in {
		crabs[i] = util.AtoIorEXIT(crab)
	}
	for _, start := range crabs {
		if start < min {
			min = start
		}
		if start > max {
			max = start
		}
	}
	fuel := math.MaxUint32
	for i := min; i <= max; i++ {
		usedFuel := 0
		for _, crab := range crabs {
			usedFuel += util.Diff(crab, i)
		}
		if usedFuel < fuel {
			fuel = usedFuel
		}
	}
	return fuel
}
