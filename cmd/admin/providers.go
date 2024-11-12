package admin

import (
	"github.com/yasindce1998/skill-marketplace/api/controllers"

	"github.com/spf13/cobra"
)

var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "Manage providers",
	Run: func(cmd *cobra.Command, args []string) {
		controllers.GetProviders(c)
	},
}

func init() {
	rootCmd.AddCommand(providersCmd)
}