package day4

import (
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 4, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day4, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	in := input.GetStringSlice()
	firstLine := strings.Split(in[0], ",")
	drag := make([]int, len(firstLine))
	for i := range drag {
		drag[i] = util.AtoIorEXIT(firstLine[i])
	}
	boards := fillBoards(in[2:])
	results := make(chan result)
	var wg sync.WaitGroup

	wg.Add(len(boards))
	for i := range boards {
		go boards[i].tryToComplete(&wg, drag, results)
	}
	var output result
	for i := 0; i < len(boards); i++ {
		tmp := <-results
		if !output.win || tmp.moves < output.moves {
			output = tmp
		}
	}
	wg.Wait()
	close(results)
	return output.value
}
