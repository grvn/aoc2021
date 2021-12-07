package day5

import (
	"regexp"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part1() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "Day 5, Problem 1",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute1(util.FromFile()),
			})).Info("Day5, Part 1")
		},
	}
}

func execute1(input *util.Input) int {
	lines := input.GetLines()
	// https://stackoverflow.com/questions/32987215/find-numbers-in-string-using-golang-regexp
	regex := regexp.MustCompile("[0-9]+")
	b := board{}
	for line := range lines {
		match := regex.FindAllString(line, -1)
		x1 := util.AtoIorEXIT(match[0])
		x2 := util.AtoIorEXIT(match[2])
		y1 := util.AtoIorEXIT(match[1])
		y2 := util.AtoIorEXIT(match[3])
		b.markOnBoard(x1, x2, y1, y2)
	}
	return b.leastDanger()
}

func (board board) markOnBoard(x1 int, x2 int, y1 int, y2 int) {
	var from, to int
	if x1 == x2 {
		if y1 > y2 {
			to = y1
			from = y2
		} else {
			to = y2
			from = y1
		}
		for from <= to {
			if _, ok := board[from]; !ok {
				board[from] = map[int]int{}
			}
			board[from][x1]++
			from++
		}
	} else if y1 == y2 {
		if x1 > x2 {
			to = x1
			from = x2
		} else {
			to = x2
			from = x1
		}
		if _, ok := board[y1]; !ok {
			board[y1] = map[int]int{}
		}
		for from <= to {
			board[y1][from]++
			from++
		}
	}
}
