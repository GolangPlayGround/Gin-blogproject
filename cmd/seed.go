package cmd

import (
	"Gin-blogproject/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seeCmd)
}

var seeCmd = &cobra.Command{
	Use:   "seed",
	Short: "Database seeder",
	Long:  "Database seeder",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	bootstrap.Seed()
}
