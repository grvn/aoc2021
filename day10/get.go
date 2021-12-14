package day10

import "github.com/spf13/cobra"

func AddCommand(cmd *cobra.Command) {
	day := &cobra.Command{
		Use:   "10",
		Short: "Day 10",
	}

	day.AddCommand(Part1())
	day.AddCommand(Part2())

	cmd.AddCommand(day)
}
