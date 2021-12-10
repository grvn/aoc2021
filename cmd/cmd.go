package cmd

import (
	"time"

	"github.com/grvn/aoc2021/day1"
	"github.com/grvn/aoc2021/day2"
	"github.com/grvn/aoc2021/day3"
	"github.com/grvn/aoc2021/day4"
	"github.com/grvn/aoc2021/day5"
	"github.com/grvn/aoc2021/day6"
	"github.com/grvn/aoc2021/day7"
	"github.com/grvn/aoc2021/day8"
	"github.com/grvn/aoc2021/day9"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

func NewCommand() *cobra.Command {
	var (
		start time.Time
	)
	result := &cobra.Command{
		Use:     "aoc2021",
		Short:   "Advent of Code 2021 Solutions",
		Example: "go run main.go 1 1 -i ./day1/input.txt",
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			start = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			log.WithFields((log.Fields{
				"Took": time.Since(start),
			})).Info("Time")
		},
	}

	day1.AddCommand(result)
	day2.AddCommand(result)
	day3.AddCommand(result)
	day4.AddCommand(result)
	day5.AddCommand(result)
	day6.AddCommand(result)
	day7.AddCommand(result)
	day8.AddCommand(result)
	day9.AddCommand(result)

	flags := result.PersistentFlags()

	flags.StringP("input", "i", "", "Input File to read. If not specified, assumes ./challenge/dayN/input.txt for the currently running challenge")

	_ = viper.BindPFlags(flags)

	return result
}
