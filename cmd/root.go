package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pbdb",
	Short: "A key/value database inspired by ch 3 of DDIA",
	Long: `A key/value database inspired by chapter 3 of Designing
Data-Intensive Applications by Martin Kleppmann.`,
}

// Execute ...
func Execute() {
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
