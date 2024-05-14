package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func Run() {
	data, err := os.ReadFile("templates.yaml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var d Data
	if err := yaml.Unmarshal(data, &d); err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		os.Exit(1)
	}

	for key, config := range d {
		fmt.Printf("Key: %s\n", key)
		for _, template := range config.Templates {
			fmt.Printf("  Template: %s\n", template.Template)
			fmt.Printf("  Path: %s\n", template.Path)
		}
		fmt.Println()
	}
}
