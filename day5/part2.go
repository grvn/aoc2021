package day5

import (
	"regexp"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 5, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day5, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
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
		b.markDiagOnBoard(x1, x2, y1, y2)
	}
	return b.leastDanger()
}

func (board board) markDiagOnBoard(x1 int, x2 int, y1 int, y2 int) {
	dx, dy := 0, 0
	if x1 > x2 {
		dx = -1
	} else if x1 < x2 {
		dx = 1
	}
	if y1 > y2 {
		dy = -1
	} else if y1 < y2 {
		dy = 1
	}
	for {
		if _, ok := board[y1]; !ok {
			board[y1] = map[int]int{}
		}
		board[y1][x1]++
		if x1 == x2 && y1 == y2 {
			break
		}
		y1 += dy
		x1 += dx
	}
}
