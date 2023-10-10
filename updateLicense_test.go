package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateLicense(t *testing.T) {
	testDir := t.TempDir()

	licensePath := createTempLicenseFile(t, testDir)
	defer os.Remove(licensePath)

	createTempTestFiles(t, testDir)

	// Dry-run test
	logOutput := captureLogOutput(func() {
		err := updateLicense(testDir, true, licensePath)
		require.NoError(t, err)
	})

	assert.Contains(t, logOutput, "Dry-run: would update license in")

	// Act
	err := updateLicense(testDir, false, licensePath)
	require.NoError(t, err)

	assertLicensePrependedToFiles(t, testDir, licensePath)
}

func createTempLicenseFile(t *testing.T, testDir string) string {
	licensePath := testDir + "/license.txt"
	licenseText := "This is the license text."
	err := os.WriteFile(licensePath, []byte(licenseText), 0644)
	require.NoError(t, err)
	return licensePath
}

func createTempTestFiles(t *testing.T, testDir string) {
	testFilePath1 := testDir + "/file1.go"
	testContent1 := "This is some test content.\n"
	err := os.WriteFile(testFilePath1, []byte(testContent1), 0644)
	require.NoError(t, err)

	testFilePath2 := testDir + "/file2.txt"
	testContent2 := "This is not a Go file.\n"
	err = os.WriteFile(testFilePath2, []byte(testContent2), 0644)
	require.NoError(t, err)
}

// assertLicensePrependedToFiles - ensures that the license was prepended to the .go files
func assertLicensePrependedToFiles(t *testing.T, testDir, licensePath string) {
	testFiles, err := os.ReadDir(testDir)
	require.NoError(t, err)

	for _, fileInfo := range testFiles {
		if !fileInfo.IsDir() && fileInfo.Name() != "license.txt" {
			filePath := testDir + "/" + fileInfo.Name()
			fileContent, err := os.ReadFile(filePath)
			require.NoError(t, err)

			licenseContent, err := os.ReadFile(licensePath)
			require.NoError(t, err)

			if strings.HasSuffix(fileInfo.Name(), ".go") {
				// Ensure that the license is prepended to the file
				assert.Contains(t, string(fileContent), string(licenseContent))
			} else {
				// Ensure that the license is not prepended to the file because it is not a Go file
				assert.NotContains(t, string(fileContent), string(licenseContent))
			}
		}
	}
}

func captureLogOutput(fn func()) string {
	logOutput := ""
	log.SetOutput(logWriter{&logOutput})
	fn()
	log.SetOutput(os.Stdout)
	return logOutput
}

type logWriter struct {
	output *string
}

func (w logWriter) Write(p []byte) (n int, err error) {
	*w.output += string(p)
	return len(p), nil
}

func TestWalker(t *testing.T) {
	testDir := t.TempDir()
	createTempTestFiles(t, testDir)

	// Act dry
	logOutput := captureLogOutput(func() {
		err := filepath.Walk(testDir, walker(true, "license text"))
		require.NoError(t, err)
	})

	assert.Contains(t, logOutput, testDir)

	// Act
	logOutput = captureLogOutput(func() {
		err := filepath.Walk(testDir, walker(false, "license text"))
		require.NoError(t, err)
	})

	assert.Contains(t, logOutput, "Processed "+testDir+"/file1.go")
}
