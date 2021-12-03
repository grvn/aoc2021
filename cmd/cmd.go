package cmd

import (
	"time"

	"github.com/grvn/aoc2021/day1"
	"github.com/grvn/aoc2021/day2"
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

	flags := result.PersistentFlags()

	flags.StringP("input", "i", "", "Input File to read. If not specified, assumes ./challenge/dayN/input.txt for the currently running challenge")

	_ = viper.BindPFlags(flags)

	return result
}
