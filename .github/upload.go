package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	srcDir := "build"
	destZip := "build-" + time.Now().Format("20060102150405") + ".zip"

	fmt.Printf("📦 Zipping folder '%s' to '%s'...\n", srcDir, destZip)

	err := zipFolder(srcDir, destZip)
	if err != nil {
		log.Fatalf("❌ Failed to zip folder: %v", err)
	}

	fmt.Printf("✅ Successfully created %s\n", destZip)
}

// zipFolder compresses an entire directory into a zip file
func zipFolder(source, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories (we’ll include them implicitly)
		if info.IsDir() {
			return nil
		}

		// Create the zip header
		relPath, err := filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath
		header.Method = zip.Deflate // compression

		// Create writer for the file
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Open source file
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Copy file contents into zip
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
