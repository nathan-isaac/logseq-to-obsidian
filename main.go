package main

import (
	"errors"
	"flag"
	"io/fs"
	"log"
	"os"
)

func main() {
	logseqDir := flag.String("logseq-dir", "", "Path to the Logseq directory")
	obsidianDir := flag.String("obsidian-dir", "", "Path to the Obsidian directory")

	flag.Parse()

	if *logseqDir == "" || *obsidianDir == "" {
		log.Fatalf("Logseq and Obsidian directories are required")
	}

	logseqFS, err := NewLogseqFS(os.DirFS(*logseqDir))
	if err != nil {
		log.Fatalf("Error creating Logseq file system: %v", err)
	}
	obsidianFS, err := NewObsidianFS(os.DirFS(*obsidianDir))
	if err != nil {
		log.Fatalf("Error creating Obsidian file system: %v", err)
	}

	app := NewApp(logseqFS, obsidianFS)
	err = app.Sync()

	if err != nil {
		log.Fatalf("Error syncing: %v", err)
	}
}

type LogseqFS struct {
	System fs.FS
}

func NewLogseqFS(system fs.FS) (*LogseqFS, error) {
	folderInfo, err := fs.Stat(system, "logseq")

	if err != nil {
		return nil, err
	}

	if !folderInfo.IsDir() {
		return nil, errors.New("logseq is not a directory")
	}

	return &LogseqFS{
		System: system,
	}, nil
}

type ObsidianFS struct {
	System fs.FS
}

func NewObsidianFS(system fs.FS) (*ObsidianFS, error) {
	folderInfo, err := fs.Stat(system, ".obsidian")

	if err != nil {
		return nil, err
	}

	if !folderInfo.IsDir() {
		return nil, errors.New("obs is not a directory")
	}

	return &ObsidianFS{
		System: system,
	}, nil
}

type App struct {
	LogseqFS   *LogseqFS
	ObsidianFS *ObsidianFS
}

func NewApp(logseqFS *LogseqFS, obsidianFS *ObsidianFS) *App {
	return &App{
		LogseqFS:   logseqFS,
		ObsidianFS: obsidianFS,
	}
}

func (a *App) Sync() error {
	err := a.SyncAssets()
	if err != nil {
		return err
	}
	a.SyncJournal()
	a.SyncPages()

	return nil
}

func (a *App) SyncAssets() error {
	log.Println("Syncing assets")
	entries, err := fs.ReadDir(a.LogseqFS.System, "assets")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		log.Printf("Entry: %s", entry.Name())
	}

	return nil
}

func (a *App) SyncJournal() {
	log.Println("Syncing journal")

}

func (a *App) SyncPages() {
	log.Println("Syncing pages")
}
