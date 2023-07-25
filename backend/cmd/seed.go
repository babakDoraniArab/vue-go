/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/babakDoraniArab/vue-go/pkg/bootstrap"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed the database with some data",
	Long:  "seed the database with some data",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	bootstrap.Seed()
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
