package main

import (
	"os"
)

// In-memory prependLicense function
func prependLicense(filePath string, licenseText string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("/*\n" + licenseText + "\n*/\n\n")
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}
