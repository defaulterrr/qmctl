package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qmctl",
	Short: "qmctl is a tool to orchestrate your virtual machines",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var testCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test launched")
		fmt.Println(verbose)
	},
}

var verbose bool

func Execute() error {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVarP(&DeployFile, "file", "f", "", "Path to the deploy description (Should be a YAML config)")
	return rootCmd.Execute()
}
