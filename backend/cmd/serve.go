/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/babakDoraniArab/vue-go/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve app on dev server",
	Long:  `Application will be served on host and port defined in config.yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {

		//let's serve the app
		serve()
	},
}

func serve() {
	bootstrap.Serve()
}
