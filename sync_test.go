package main

import (
	"io/fs"
	"log"
	"testing"
	"testing/fstest"
)

func TestLogseqFS(t *testing.T) {
	t.Run("missing logseq directory", func(t *testing.T) {
		testFs := fstest.MapFS{}

		_, err := NewLogseqFS(testFs)

		if err == nil {
			t.Errorf("Expected error creating Logseq file system")
		}
	})

	t.Run("logseq directory exists", func(t *testing.T) {
		testFs := fstest.MapFS{
			"logseq": {Mode: fs.ModeDir},
		}

		_, err := NewLogseqFS(testFs)

		if err != nil {
			t.Errorf("Error creating Logseq file system: %v", err)
		}
	})
}

func TestObsidianFS(t *testing.T) {
	t.Run("missing obsidian directory", func(t *testing.T) {
		testFs := fstest.MapFS{}

		_, err := NewObsidianFS(testFs)

		if err == nil {
			t.Errorf("Expected error creating Obsidian file system")
		}
	})

	t.Run("obsidian directory exists", func(t *testing.T) {
		testFs := fstest.MapFS{
			".obsidian": {Mode: fs.ModeDir},
		}

		_, err := NewObsidianFS(testFs)

		if err != nil {
			t.Errorf("Error creating Obsidian file system: %v", err)
		}
	})
}

func TestExample(t *testing.T) {
	testFs := fstest.MapFS{
		// assets
		"assets/image.png":    {Data: []byte("hi")},
		"assets/image.jpg":    {Data: []byte("hi")},
		"assets/image.jpeg":   {Data: []byte("hi")},
		"assets/document.pdf": {Data: []byte("hi")},
		// pages
		//"pages/hello world.md":   {Data: []byte("hi")},
		//"journals/2024_12_28.md": {Data: []byte("hi")},
	}

	//srcPath := "obs/hello world.md"
	//destPath := "log/hello world.md"
	//

	// check if logseq directory exists
	// check if .obsidian directory exists
	err := fs.WalkDir(testFs, "obs", func(path string, d fs.DirEntry, err error) error {
		// check if path exists in logseq directory
		// if not, copy the file

		log.Printf("Path: %s", path)
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}

	//entries, err := fs.ReadDir("obs")
	//
	//if err != nil {
	//	t.Errorf("Error reading directory: %v", err)
	//}
	//
	//for _, entry := range entries {
	//	if entry.IsDir() {
	//		subEntries, err := fs.ReadDir("obs/" + entry.Name())
	//		if err != nil {
	//			t.Errorf("Error reading directory: %v", err)
	//		}
	//		for _, subEntry := range subEntries {
	//			if !subEntry.IsDir() {
	//				t.Logf("Entry: %s", subEntry.Name())
	//			}
	//		}
	//	} else {
	//		t.Logf("Entry: %s", entry.Name())
	//	}
	//}

	//app := NewApp()
	//app.Sync()
}
