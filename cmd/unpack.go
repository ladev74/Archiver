package cmd

import "github.com/spf13/cobra"

var unpackCmd = &cobra.Command{
	Use:   "pack",
	Short: "Unpack file",
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}
