package day1

import "github.com/spf13/cobra"

func AddCommand(cmd *cobra.Command) {
	day := &cobra.Command{
		Use:   "1",
		Short: "Day 1",
	}

	day.AddCommand(Part1())
	day.AddCommand(Part2())

	cmd.AddCommand(day)
}
