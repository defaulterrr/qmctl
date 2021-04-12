package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qmctl",
	Short: "qmctl is a tool to orchestrate your virtual machines",
}

var testCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test launched")
		fmt.Println(Verbose)
	},
}

var DeployFile string
var Verbose bool
var Testing bool

func Execute() error {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Testing, "testing", "", false, "[TEST] intended for test purposes, allows for some workarounds to be activated when full environment is not available")
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(flushCmd)
	applyCmd.Flags().StringVarP(&DeployFile, "file", "f", "", "Path to the deploy description (Should be a YAML config)")
	return rootCmd.Execute()
}
