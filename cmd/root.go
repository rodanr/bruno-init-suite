package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var outputDir string

var rootCmd = &cobra.Command{
	Use:   "bruis", // short form for bruno init suite
	Short: "CLI tool for generating bruno docs with custom configurations",
	Long:  `A CLI tool to automate the creation of Bruno docs and integration with third part services, including scripts and configurations`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Oops. An error occurred while executing Bruno Init Suite '%s'\n", err)
		os.Exit(1)
	}
}

func init() {
	// Global flag for output directory
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "Output directory for generated Bruno docs")
}
