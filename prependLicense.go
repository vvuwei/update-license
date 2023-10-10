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
	commentedLicenseLines := make([]string, len(licenseLines)+2)
	commentedLicenseLines[0] = "/*"
	for i := 1; i < len(licenseLines)+1; i++ {
		commentedLicenseLines[i] = licenseLines[i-1]
	}
	commentedLicenseLines[len(licenseLines)+1] = "*/"

	content := string(fileContent)

	if hasLicense(content, licenseLines) {
		return nil
	}

	newContent := []byte(strings.Join(append(commentedLicenseLines, "", content), "\n"))
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
