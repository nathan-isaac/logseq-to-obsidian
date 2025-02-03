package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	logseqDir := flag.String("logseq-dir", "", "Path to the Logseq directory")
	obsidianDir := flag.String("obsidian-dir", "", "Path to the Obsidian directory")

	flag.Parse()

	absLogseqDir, err := filepath.Abs(*logseqDir)

	if err != nil {
		log.Fatalf("Error geting absolute path for Logseq directory: %v", err)
	}

	absObsidianDir, err := filepath.Abs(*obsidianDir)

	if err != nil {
		log.Fatalf("Error geting absolute path for Obsidian directory: %v", err)
	}

	fmt.Printf("Logseq Directory: %s\n", absLogseqDir)
	fmt.Printf("Obsidian Directory: %s\n", absObsidianDir)
}
