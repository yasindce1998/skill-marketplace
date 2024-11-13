package cmd

import (
	"fmt"
	"github.com/yasindce1998/skill-marketplace/api/controllers"
	"time"

	"github.com/spf13/cobra"
)


var statsCmd = &cobra.Command{
	Use:   "stats <start_date> <end_date>",
	Short: "Get periodic statistics",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		startDate, err := time.Parse("2006-01-02", args[0])
		if err != nil {
			fmt.Printf("Invalid start date: %v\n", err)
			return
		}

		endDate, err := time.Parse("2006-01-02", args[1])
		if err != nil {
			fmt.Printf("Invalid end date: %v\n", err)
			return
		}

		stats := controllers.GetPeriodicStats(startDate, endDate)
		fmt.Printf("Total tasks: %d\n", stats.TotalTasks)
		fmt.Printf("Completed tasks: %d\n", stats.CompletedTasks)
		fmt.Printf("Rejected tasks: %d\n", stats.RejectedTasks)
		fmt.Printf("Average provider success ratio: %.2f%%\n", stats.AvgProviderSuccessRatio)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
