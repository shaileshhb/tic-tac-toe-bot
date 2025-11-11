package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run check_build_file.go <path-to-file>")
	}

	filePath := os.Args[1]

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Fatalf("❌ File not found: %s", filePath)
	} else if err != nil {
		log.Fatalf("⚠️ Error accessing file: %v", err)
	}

	fmt.Printf("✅ File found: %s\n", filePath)
	fmt.Printf("📦 Size: %d bytes\n", info.Size())
	fmt.Printf("🕒 Last modified: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
}
