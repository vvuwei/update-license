package main

import (
	"os"
	"testing"
)

func BenchmarkPrependLicense(b *testing.B) {
	tmpDir := b.TempDir()
	testFilePath := tmpDir + "/test.go"
	testContent := source
	err := os.WriteFile(testFilePath, []byte(testContent), 0644)
	if err != nil {
		b.Fatalf("Error creating test file: %v", err)
	}
	licenseText := license

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := prependLicense(testFilePath, licenseText)
		if err != nil {
			b.Fatalf("Error calling prependLicenseInMemoryWithCheck: %v", err)
		}
	}
}
