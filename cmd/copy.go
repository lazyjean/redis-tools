/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lazyjean/redis-tools/app"
	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: app.Copy,
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringP("src-host", "", "127.0.0.1", "src host")
	copyCmd.Flags().IntP("src-port", "", 6379, "src port")
	copyCmd.Flags().StringP("src-password", "", "", "src password")
	copyCmd.Flags().StringP("dst-host", "", "127.0.0.1", "dst host")
	copyCmd.Flags().IntP("dst-port", "", 6379, "dst port")
	copyCmd.Flags().StringP("dst-password", "", "", "dst password")
	copyCmd.Flags().StringP("key-prefix", "k", "", "redis key")
	copyCmd.Flags().BoolP("zset", "z", false, "zSet")
}
