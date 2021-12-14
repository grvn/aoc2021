package day10

import (
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

var acpoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 10, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day10, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
	scores := []int{}
exit:
	for line := range input.GetLines() {
		linescore := 0
		queue := util.NewQueue()
		for _, char := range line {
			if close, ok := closes[char]; ok {
				queue.Push(close)
				continue
			}
			top, ok := queue.Peek()
			if !ok {
				continue exit
			}

			if char == top.(rune) {
				queue.Pop()
				continue
			}
			continue exit
		}

		for queue.Length() > 0 {
			linescore *= 5
			next, _ := queue.Pop()
			linescore += acpoints[next.(rune)]
		}
		scores = append(scores, linescore)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
