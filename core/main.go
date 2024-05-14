package core

import (
	"embed"
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"gopkg.in/yaml.v2"
)

//go:embed schema.cue
var schemaFS embed.FS

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

func Run2() {
	ctx := cuecontext.New()

	schemaBytes, err := schemaFS.ReadFile("schema.cue")
	if err != nil {
		fmt.Println("Error reading schema file:", err)
		os.Exit(1)
	}
	schema := ctx.CompileBytes(schemaBytes, cue.Filename("schema.cue"))
	if err := schema.Err(); err != nil {
		fmt.Println("Error compiling schema:", err)
		os.Exit(1)
	}

	yamlBytes, err := os.ReadFile("templates.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		os.Exit(1)
	}

	var d Data
	if err := yaml.Unmarshal(yamlBytes, &d); err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		os.Exit(1)
	}

	data := ctx.Encode(d)
	if err := data.Err(); err != nil {
		fmt.Println("Error encoding data:", err)
		os.Exit(1)
	}

	v := schema.Unify(data)
	if err := v.Err(); err != nil {
		fmt.Println("Error unifying schema and data:", err)
		os.Exit(1)
	}
	if err := v.Validate(); err != nil {
		fmt.Println("Validation failed:", err)
		os.Exit(1)
	}

	fmt.Println("Validation succeeded")
}
