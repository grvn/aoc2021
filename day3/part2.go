package day3

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grvn/aoc2021/util"
)

const (
	MOST  = true
	LEAST = false
)

func Part2() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "Day 3, Problem 2",
		Run: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Answer": execute2(util.FromFile()),
			})).Info("Day3, Part 2")
		},
	}
}

func execute2(input *util.Input) int {
	lines := input.GetStringSlice()
	col := len(lines[0])
	remaining := make([]string, len(lines))
	o2, co2 := 0, 0
	copy(remaining, lines)

	for i := 0; i < col; i++ {
		value := commonValue(remaining, i, MOST)
		remains := make([]string, 0, len(remaining))
		for _, n := range remaining {
			if rune(n[i]) == value {
				remains = append(remains, n)
			}
		}
		if len(remains) == 1 {
			v, _ := strconv.ParseInt(remains[0], 2, 32)
			o2 = int(v)
		}
		remaining = remains
	}

	remaining = make([]string, len(lines))
	copy(remaining, lines)
	for i := 0; i < col; i++ {
		value := commonValue(remaining, i, LEAST)

		remains := make([]string, 0, len(remaining))
		for _, n := range remaining {
			if rune(n[i]) == value {
				remains = append(remains, n)
			}
		}
		if len(remains) == 1 {
			v, _ := strconv.ParseInt(remains[0], 2, 32)
			co2 = int(v)
		}
		remaining = remains
	}

	return o2 * co2
}

func commonValue(lines []string, bit int, most bool) rune {
	zeros, ones := 0, 0
	for n := 0; n < len(lines); n++ {
		switch lines[n][bit] {
		case '0':
			zeros++
		case '1':
			ones++
		}
	}
	if zeros > ones {
		if most {
			return '0'
		} else {
			return '1'
		}
	} else {
		if most {
			return '1'
		} else {
			return '0'
		}
	}
}
