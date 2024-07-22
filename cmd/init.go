package cmd

import (
	"bruno-init-suite/internal/initializer"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

var useCognito bool
var projectName string
var baseUrl string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Bruno project",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Absolute path for outputDir
		absOutputDir, err := filepath.Abs(outputDir)
		if err != nil {
			return fmt.Errorf("failed to determine absolute path: %w", err)
		}

		err = initializer.Initialize(absOutputDir, projectName, baseUrl, useCognito)
		if err != nil {
			return err
		}

		fmt.Println("Files generated successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// Add flags to the init command
	// The flag useCognito is used to determine if the project will use AWS Cognito for authentication
	initCmd.Flags().BoolVarP(&useCognito, "cognito", "c", false, "Use AWS Cognito for authentication")
	// The flag for name is used to set the project name
	initCmd.Flags().StringVarP(&projectName, "name", "n", "bruno-docs", "Name of the project")
	// The flag for baseUrl is used to set the base URL for the project
	initCmd.Flags().StringVarP(&baseUrl, "baseUrl", "b", "http://localhost:8080", "Base URL for the project")
}
