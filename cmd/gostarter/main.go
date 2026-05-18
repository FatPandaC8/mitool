package main

import (
	"gostarter"
	"flag"
	"fmt"
	"gostarter/internal"
	"os"
	"path/filepath"
)

const Version = "v0.0.1"

func main() {
	if len(os.Args) < 2 {
		printGlobalUsage()
		return
	}

	switch os.Args[1] {
	case "--version":
		fmt.Printf("gostarter version %s\n", Version)
		return

	case "init":
		fmt.Println("Initializing global gostarter configuration...")
		configDir, err := runGlobalSetup()
		if err != nil {
			fmt.Printf("Initialization failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully initialized global templates folder at: %s\n", filepath.Join(configDir, "templates"))
		return

	case "new":
		newCmd := flag.NewFlagSet("new", flag.ExitOnError)

		var lang string
		var templateName string

		newCmd.StringVar(&lang, "lang", "go", "Programming language for the project")
		newCmd.StringVar(&templateName, "template", "basic", "Template name found in the templates/ directory")

		newCmd.Usage = func() {
			fmt.Println("Usage: gostarter new [project-name] [flags]")
			fmt.Println("\nFlags:")
			newCmd.PrintDefaults()
		}

		if len(os.Args) < 3 {
			newCmd.Usage()
			return
		}

		// The project name must follow right after 'new'
		projectName := os.Args[2]

		// Parse the flags trailing the project name
		_ = newCmd.Parse(os.Args[3:])

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			os.Exit(1)
		}

		templatePath := filepath.Join(homeDir, ".config", "gostarter", "templates", lang, templateName+".yaml")

		fmt.Printf("Loading template from: %s\n", templatePath)
		
		// 1. Load template using your existing internal code
		tmpl, err := internal.LoadTemplate(templatePath)
		if err != nil {
			fmt.Printf("Error loading template: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Generating project structure for '%s'...\n", projectName)

		// 2. Build the directory tree using your existing internal code
		err = internal.CreateProject(projectName, tmpl)
		if err != nil {
			fmt.Printf("Error generating project: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Project '%s' generated successfully!\n", projectName)

	default:
		fmt.Printf("Unknown command: %q\n", os.Args[1])
		printGlobalUsage()
		os.Exit(1)
	}
}

func printGlobalUsage() {
	fmt.Println("Usage:")
	fmt.Println("  gostarter <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  new       Create a new project using a template blueprint")
	fmt.Println("  version   Show CLI version")
}

func runGlobalSetup() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "gostarter")
	goTemplateDir := filepath.Join(configDir, "templates", "go")

	// Force-build the entire workspace block directory hierarchy
	err = os.MkdirAll(goTemplateDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create global templates directory: %w", err)
	}

	// Always ensure your standard basic template gets written or updated
	defaultTargetFile := filepath.Join(goTemplateDir, "basic.yaml")
	err = os.WriteFile(defaultTargetFile, gostarter.DefaultGoBasicTemplate, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write default basic.yaml template structure: %w", err)
	}

	return configDir, nil
}