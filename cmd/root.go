/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "redis-tools",
	Short: "a collection of tools for redis",
	Long:  `a collection of tools for redis, include copy, move`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(moveCmd)

	rootCmd.PersistentFlags().StringP("src-host", "", "127.0.0.1", "src host")
	rootCmd.PersistentFlags().IntP("src-port", "", 6379, "src port")
	rootCmd.PersistentFlags().StringP("src-password", "", "", "src password")
	rootCmd.PersistentFlags().StringP("dst-host", "", "127.0.0.1", "dst host")
	rootCmd.PersistentFlags().IntP("dst-port", "", 6379, "dst port")
	rootCmd.PersistentFlags().StringP("dst-password", "", "", "dst password")
	rootCmd.PersistentFlags().StringP("key-pattern", "k", "", "key pattern")
	rootCmd.PersistentFlags().BoolP("zset", "z", false, "zSet")
}
