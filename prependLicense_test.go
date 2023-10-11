package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const content = "This is some test content.\n"
const licenseTxt = "This is the license text."

func Test_PrependInMemory(t *testing.T) {
	tmpDir := t.TempDir()
	testFilePath := tmpDir + "/test.go"
	testContent := content
	err := os.WriteFile(testFilePath, []byte(testContent), 0644)
	require.NoError(t, err)

	licenseText := licenseTxt

	// Act
	err = prependLicense(testFilePath, licenseText)
	require.NoError(t, err)

	modifiedContent, err := os.ReadFile(testFilePath)
	require.NoError(t, err)

	// Check if the modified content starts with the license text
	expectedContent := "// " + licenseText + "\n\n" + testContent
	assert.Equal(t, expectedContent, string(modifiedContent))

	// Check if the original content is preserved
	assert.NotEqual(t, testContent, string(modifiedContent))
}

func Test_PrependInMemory_UnhappyPath(t *testing.T) {
	tmpDir := t.TempDir()
	testFilePath := tmpDir + "/test.go"
	testContent := content
	err := os.WriteFile(testFilePath, []byte(testContent), 0644)
	require.NoError(t, err)

	// Create a file with read-only permissions to simulate an error
	err = os.WriteFile(testFilePath+".readonly", []byte(testContent), 0444)
	require.NoError(t, err)

	licenseText := licenseTxt

	// Act
	err = prependLicense(testFilePath+".readonly", licenseText)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "permission denied")

	// Ensure that the original file content is not modified
	originalContent, err := os.ReadFile(testFilePath + ".readonly")
	require.NoError(t, err)

	assert.Equal(t, testContent, string(originalContent))
}

func TestHasLicense(t *testing.T) {
	licenseLines := []string{
		"/*",
		"    This is the license text.",
		"    */",
	}

	var tests = []struct {
		name     string
		content  string
		license  []string
		expected bool
	}{
		{
			name: "License is present",
			content: `/*
    This is the license text.
    */
    package main`,
			license:  licenseLines,
			expected: true,
		},
		{
			name:     "License is not present",
			content:  `package main`,
			license:  licenseLines,
			expected: false,
		},
		{
			name:     "Content is shorter than the license",
			content:  `/* This is a comment. */ package main`,
			license:  licenseLines,
			expected: false,
		},
		{
			name:     "Empty license, should always return true",
			content:  `package main`,
			license:  []string{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, hasLicense(tt.content, tt.license))
		})
	}
}

// Run the tests in parallel
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func BenchmarkPrependLicense(b *testing.B) {
	tmpDir := b.TempDir()
	testFilePath := tmpDir + "/test.go"
	testContent := source
	err := os.WriteFile(testFilePath, []byte(testContent), 0644)
	if err != nil {
		b.Fatalf("Error creating test file: %v", err)
	}
	licenseText := licenseTxt

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := prependLicense(testFilePath, licenseText)
		if err != nil {
			b.Fatalf("Error calling prependLicenseInMemoryWithCheck: %v", err)
		}
	}
}
