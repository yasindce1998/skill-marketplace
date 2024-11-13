package cmd

import (
	"fmt"
	"github.com/yasindce1998/skill-marketplace/api/controllers"

	"github.com/spf13/cobra"
)

var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "Get the number of providers",
	Run: func(cmd *cobra.Command, args []string) {
		providerCount := controllers.GetProviderCount()
		fmt.Printf("Number of providers: %d\n", providerCount)
	},
}

func init() {
	rootCmd.AddCommand(providersCmd)
}