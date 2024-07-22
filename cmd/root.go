package cmd

import (
	"bruno-init-suite/internal/version"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	outputDir   string
	showVersion bool
)

var rootCmd = &cobra.Command{
	Use:     "bruno-init-suite",
	Aliases: []string{"bruis"}, // short form for bruno init suite
	Short:   "CLI tool for generating bruno docs with custom configurations",
	Long:    `A CLI tool to automate the creation of Bruno docs and integration with third part services, including scripts and configurations`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("bruno-init-suite version %s\n", version.Version)
			os.Exit(0)
		}
	},
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
	// Global flag for version
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show version of Bruno Init Suite")
}
