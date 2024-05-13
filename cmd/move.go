/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lazyjean/redis-tools/app"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: app.Move,
}

func init() {
	rootCmd.AddCommand(moveCmd)
	// moveCmd.Flags().StringP("src-host", "", "127.0.0.1", "src host")
	// moveCmd.Flags().IntP("src-port", "", 6379, "src port")
	// moveCmd.Flags().StringP("src-password", "", "", "src password")
	// moveCmd.Flags().StringP("dst-host", "", "127.0.0.1", "dst host")
	// moveCmd.Flags().IntP("dst-port", "", 6379, "dst port")
	// moveCmd.Flags().StringP("dst-password", "", "", "dst password")
	// moveCmd.Flags().StringP("key-pattern", "k", "", "key pattern")
	// moveCmd.Flags().BoolP("zset", "z", false, "zSet")
}
