/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/babakDoraniArab/vue-go/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database tables",
	Long: `migrate database tables
for more information about gorm migration https://gorm.io/docs/migration.html`,
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	bootstrap.Migrate()
}
