package day8

import "github.com/spf13/cobra"

func AddCommand(cmd *cobra.Command) {
	day := &cobra.Command{
		Use:   "8",
		Short: "Day 8",
	}

	day.AddCommand(Part1())
	day.AddCommand(Part2())

	cmd.AddCommand(day)
}
