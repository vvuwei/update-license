package main

import (
	"flag"
	"log"
	"os"
)

var (
	flagPath = flag.String(
		"path",
		".",
		"Path to the folder containing the files to update")
	flagDryRun = flag.Bool(
		"dry",
		false,
		"Do not edit files and just print out what would be edited")
	flagLicense = flag.String(
		"license",
		"",
		"Path to the license file")
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.SetPrefix("")

	flag.Parse()

	if err := updateLicense(*flagPath, *flagDryRun, *flagLicense); err != nil {
		log.Fatal(err)
	}
}
