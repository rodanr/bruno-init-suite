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
//go:embed templates/*
var templateFS embed.FS

type templateData struct {
	ProjectName string
	BaseURL     string
	BaseURLType string
	UseCognito  bool
}

func Initialize(outputDir, projectName string, useCognito bool) error {
	data := templateData{
		ProjectName: projectName,
	}

	dirs := []string{
		outputDir,
		filepath.Join(outputDir, "environments"),
		filepath.Join(outputDir, "lib"),
	}

	if useCognito {
		if err := copyFile("templates/cognito-auth.js", filepath.Join(outputDir, "lib", "cognito-auth.js")); err != nil {
			return err
		}
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	files := map[string]string{
		"templates/bruno_template.json":          filepath.Join(outputDir, "bruno.json"),
		"templates/collection_template.bru":      filepath.Join(outputDir, "collection.bru"),
		"templates/env_template.env":             filepath.Join(outputDir, ".env"),
		"templates/environments/environment.bru": filepath.Join(outputDir, "environments", "environment.bru"),
	}

	for tpl, dest := range files {
		if err := generateFile(tpl, dest, data); err != nil {
			return err
		}
	}

	return nil
}

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
