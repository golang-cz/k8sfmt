package k8sfmt

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func PrettifyYAML(filename string) error {
	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Transform YAML nodes
	transformNode(&root)

	// Marshal back to YAML
	output, err := yaml.Marshal(&root)
	if err != nil {
		return fmt.Errorf("failed to serialize YAML: %w", err)
	}

	// Write back to file
	if err := os.WriteFile(filename, output, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func short(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen < 3 {
		maxLen = 3
	}
	return string(runes[0:maxLen-3]) + "..."
}

func transformNode(node *yaml.Node) {
	if node == nil {
		return
	}

	switch node.Kind {
	case yaml.DocumentNode, yaml.MappingNode, yaml.SequenceNode:
		for _, child := range node.Content {
			transformNode(child)
		}

	case yaml.ScalarNode:
		log.Printf("style: %v, tag: %v, value: %v", node.Style, node.Tag, short(node.Value, 100))

		if node.Tag == "!!str" && containsNewline(node.Value) {
			log.Printf("applying block style on node.Style=%v", node.Style)
			//node.Style = yaml.LiteralStyle // Apply block scalar style
			node.Style = yaml.LiteralStyle
			node.Value = "." + node.Value
		}
	}
}

func containsNewline(s string) bool {
	return len(s) > 0 && (contains(s, "\n") || contains(s, "\r"))
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (stringIndex(s, substr) >= 0)
}

func stringIndex(s, substr string) int {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
