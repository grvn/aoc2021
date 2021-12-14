package day11

import "github.com/spf13/cobra"

func AddCommand(cmd *cobra.Command) {
	day := &cobra.Command{
		Use:   "11",
		Short: "Day 11",
	}

	day.AddCommand(Part1())
	day.AddCommand(Part2())

	cmd.AddCommand(day)
}
