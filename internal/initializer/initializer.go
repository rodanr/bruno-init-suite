package initializer

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

// Embed templates
//
//go:embed templates/*.tmpl templates/environments/*.tmpl templates/lib/*.tmpl
var templateFS embed.FS

type templateData struct {
	ProjectName string
	UseBaseURL  bool
	BaseURL     string
	BaseURLType string
	UseCognito  bool
}

// Initialize initializes a new Bruno project in the specified output directory.
func Initialize(outputDir, projectName, baseUrl string, useCognito bool) error {
	data := templateData{
		ProjectName: projectName,
		UseCognito:  useCognito,
		BaseURL:     baseUrl,
		UseBaseURL:  baseUrl != "",
		BaseURLType: "http",
	}

	dirs := []string{
		outputDir,
		filepath.Join(outputDir, "environments"),
		filepath.Join(outputDir, "lib"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	files := map[string]string{
		"templates/bruno_template.json.tmpl":          filepath.Join(outputDir, "bruno.json"),
		"templates/collection_template.bru.tmpl":      filepath.Join(outputDir, "collection.bru"),
		"templates/env_template.env.tmpl":             filepath.Join(outputDir, ".env.example"),
		"templates/environments/environment.bru.tmpl": filepath.Join(outputDir, "environments", "environment.bru"),
	}

	for tpl, dest := range files {
		if err := generateFile(tpl, dest, data); err != nil {
			return err
		}
	}

	if useCognito {
		if err := copyFile("templates/lib/cognito_auth.js.tmpl", filepath.Join(outputDir, "lib", "cognito-auth.js")); err != nil {
			return err
		}
	}

	return nil
}

// generateFile generates a file from a template file and writes it to the output path.
func generateFile(templatePath, outputPath string, data templateData) error {
	tmplContent, err := templateFS.ReadFile(templatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(tmplContent))
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := tmpl.Execute(outFile, data); err != nil {
		return err
	}

	return nil
}

// copyFile copies a file from the source path to the destination path.
func copyFile(srcPath, dstPath string) error {
	srcFile, err := templateFS.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := srcFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := dstFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
