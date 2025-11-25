package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadAllTerraformFiles reads all .tf files in a directory and its subdirectories
func ReadAllTerraformFiles(rootPath string) (string, error) {
	var allContent strings.Builder

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".tf") {
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading %s: %w", path, err)
			}
			allContent.WriteString(string(content))
			allContent.WriteString("\n")
		}
		return nil
	})

	return allContent.String(), err
}
