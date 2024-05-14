/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lazyjean/redis-tools/app"

	"github.com/spf13/cobra"
)

// unionCmd represents the union command
var unionCmd = &cobra.Command{
	Use:   "union",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: app.Union,
}

func init() {
	rootCmd.AddCommand(unionCmd)
}
