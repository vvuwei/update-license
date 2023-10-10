package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <folder_path>")
		return
	}
	folderPath := os.Args[1]

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			err := prependLicense(path, "LICENSE")
			if err != nil {
				fmt.Printf("Error processing %s: %v\n", path, err)
			} else {
				fmt.Printf("Processed %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through folder: %v\n", err)
	}
}
