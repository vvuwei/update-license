package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func updateLicense(path string, dry bool, license string) error {
	var licenseText string
	if !dry {
		licenseBytes, err := os.ReadFile(license)
		if err != nil {
			return fmt.Errorf("Error reading license file: %v\n", err)
		}

		licenseText = string(licenseBytes)
	}

	if dry {
		log.Printf("Dry-run: would update license in '%s' in files:\n", path)
	}
	err := filepath.Walk(path, walker(dry, licenseText))

	if err != nil {
		return fmt.Errorf("Error walking through folder: %v\n", err)
	}

	return nil
}

func walker(dry bool, licenseText string) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".go" {
			if dry {
				log.Printf("%s\n", path)
				return nil
			}

			err := prependLicense(path, licenseText)

			if err != nil {
				return fmt.Errorf("Error processing %s: %v\n", path, err)
			}

			log.Printf("Processed %s\n", path)
		}
		return nil
	}
}
