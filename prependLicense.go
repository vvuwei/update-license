package main

import (
	"bufio"
	"os"
	"strings"
)

// In-memory prependLicense function with license check and Go comment wrapping
func prependLicense(filePath string, licenseText string) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	licenseLines := strings.Split(licenseText, "\n")
	for i := 0; i < len(licenseLines); i++ {
		licenseLines[i] = "// " + licenseLines[i]
	}

	content := string(fileContent)

	if hasLicense(content, licenseLines) {
		return nil
	}

	newContent := []byte(strings.Join(append(licenseLines, "", content), "\n"))
	err = os.WriteFile(filePath, newContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func hasLicense(content string, licenseLines []string) bool {
	scanner := bufio.NewScanner(strings.NewReader(content))
	for i := 0; i < len(licenseLines); i++ {
		if !scanner.Scan() {
			return false
		}
		if scanner.Text() != licenseLines[i] {
			return false
		}
	}

	return true
}
