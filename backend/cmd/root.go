/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Help command",
	Long:  `This is the help command. It will show you the help for all the commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("help is called")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
