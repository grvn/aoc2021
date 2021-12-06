package day4

import "github.com/spf13/cobra"

func AddCommand(cmd *cobra.Command) {
	day := &cobra.Command{
		Use:   "4",
		Short: "Day 4",
	}

	day.AddCommand(Part1())
	day.AddCommand(Part2())

	cmd.AddCommand(day)
}
