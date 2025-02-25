package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: k8sfmt <filename> [filename2] [filename3]...")
		os.Exit(1)
	}

	for _, f := range os.Args[1:] {
		if err := formatYAML(f, f); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v: %v\n", f, err)
			os.Exit(1)
		}
	}
}

func formatYAML(fromFilename string, toFilename string) error {
	// Read the YAML file
	f, err := os.ReadFile(fromFilename)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	// Unmarshal and marshal to format YAML with indent=2.
	var node yaml.Node
	if err := yaml.Unmarshal(f, &node); err != nil {
		return fmt.Errorf("unmarshal YAML: %w", err)
	}
	if len(node.Content) != 1 {
		return fmt.Errorf("expected 1 YAML node, got %d", len(node.Content))
	}

	var formatted bytes.Buffer
	enc := yaml.NewEncoder(&formatted)
	enc.SetIndent(2)

	if err := enc.Encode(node.Content[0]); err != nil {
		return fmt.Errorf("marshal YAML: %w", err)
	}

	// Read each line
	input := bufio.NewScanner(&formatted)
	input.Buffer(nil, 10*1024*1024) // Allow reading long lines, up to 10 MB.

	outputFile, err := os.Create(toFilename)
	if err != nil {
		return fmt.Errorf("open file for writing %v: %w", toFilename, err)
	}

	writer := bufio.NewWriter(outputFile)
	for input.Scan() {
		line := input.Text()
		processedLine := processLine(line)
		_, err := writer.WriteString(processedLine + "\n")
		if err != nil {
			return fmt.Errorf("write to file %v: %w", toFilename, err)
		}
	}
	if err := input.Err(); err != nil {
		return fmt.Errorf("error reading scanner: %w", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush writer: %w", err)
	}

	return outputFile.Close()
}

func processLine(line string) string {
	if len(line) == 0 {
		// Skip empty lines
		return line
	}

	trimmed := strings.TrimSpace(line)
	if len(trimmed) == 0 || trimmed[0] == '#' {
		// Skip comments or empty lines
		return line
	}

	// Find the first colon to separate key and value
	key, value, found := strings.Cut(line, ":")
	if !found {
		// No key-value pair found
		return line
	}

	// Determine indentation level for the value
	keyIndentation := strings.IndexFunc(line, func(r rune) bool { return r != ' ' })
	indentation := strings.Repeat(" ", keyIndentation+2)

	value = strings.TrimSpace(value)

	// Find out if the value is a quoted string (double quotes only)
	if len(value) >= 2 && (value[0] == '"' && value[len(value)-1] == '"') {
		// De-escape special characters inside the string content
		unquoted, err := strconv.Unquote(value)
		if err != nil {
			panic(err) // This can't happen, since the string was created via yaml.Marshal().
		}

		if !strings.Contains(unquoted, "\n") {
			// Single line string
			return line
		}

		// Prepare the block scalar for multiline string
		lines := strings.Split(unquoted, "\n")
		for i, l := range lines {
			lines[i] = indentation + l
		}
		blockScalar := fmt.Sprintf("%s: |2\n%s", key, strings.Join(lines, "\n"))
		return blockScalar
	}

	return line
}
