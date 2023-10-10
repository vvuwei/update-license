package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PrependInMemory(t *testing.T) {
	tmpDir := t.TempDir()
	testFilePath := tmpDir + "/test.go"
	testContent := "This is some test content.\n"
	err := os.WriteFile(testFilePath, []byte(testContent), 0644)
	require.NoError(t, err)

	licenseText := "This is the license text."

	// Act
	err = prependLicense(testFilePath, licenseText)
	require.NoError(t, err)

	modifiedContent, err := os.ReadFile(testFilePath)
	require.NoError(t, err)

	// Check if the modified content starts with the license text
	expectedContent := "/*\n" + licenseText + "\n*/\n\n" + testContent
	assert.Equal(t, expectedContent, string(modifiedContent))

	// Check if the original content is preserved
	assert.NotEqual(t, testContent, string(modifiedContent))
}

func Test_PrependInMemory_UnhappyPath(t *testing.T) {
	tmpDir := t.TempDir()
	testFilePath := tmpDir + "/test.go"
	testContent := "This is some test content.\n"
	err := os.WriteFile(testFilePath, []byte(testContent), 0644)
	require.NoError(t, err)

	// Create a file with read-only permissions to simulate an error
	err = os.WriteFile(testFilePath+".readonly", []byte(testContent), 0444)
	require.NoError(t, err)

	licenseText := "This is the license text."

	// Act
	err = prependLicense(testFilePath+".readonly", licenseText)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "permission denied")

	// Ensure that the original file content is not modified
	originalContent, err := os.ReadFile(testFilePath + ".readonly")
	require.NoError(t, err)

	assert.Equal(t, testContent, string(originalContent))
}

// Run the tests in parallel
func TestMain(m *testing.M) {
	// Use -parallel flag to enable parallel testing
	// For example, "go test -parallel 4"
	// Adjust the value to control the degree of parallelism
	// You can remove the flag if you don't want parallel testing
	os.Exit(m.Run())
}
