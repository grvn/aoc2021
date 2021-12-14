package day10

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

var (
	closes = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	points = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 10, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day10, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	errorscore := 0
	for line := range input.GetLines() {
		queue := util.NewQueue()
		for _, char := range line {
			if close, ok := closes[char]; ok {
				queue.Push(close)
				continue
			}
			top, ok := queue.Peek()
			if !ok {
				errorscore += points[char]
				continue
			}

			if char == top.(rune) {
				queue.Pop()
				continue
			}
			errorscore += points[char]
			break
		}
	}
	return errorscore
}
