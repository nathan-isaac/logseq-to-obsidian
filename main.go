package main

import (
	"flag"
	"log"
	"path/filepath"
)

func main() {
	app := New()
	app.AddFlagConfig()

	log.Printf("Logseq Directory: %s\n", app.LogseqDir)
	log.Printf("Obsidian Directory: %s\n", app.ObsidianDir)

	app.Sync()
}

type App struct {
	LogseqDir   string
	ObsidianDir string
}

func New() *App {
	return &App{}
}

func (a *App) AddFlagConfig() {
	logseqDir := flag.String("logseq-dir", "", "Path to the Logseq directory")
	obsidianDir := flag.String("obsidian-dir", "", "Path to the Obsidian directory")

	flag.Parse()

	absLogseqDir, err := filepath.Abs(*logseqDir)
	if err != nil {
		log.Fatalf("Error geting absolute path for Logseq directory: %v", err)
	}
	a.LogseqDir = absLogseqDir

	absObsidianDir, err := filepath.Abs(*obsidianDir)
	if err != nil {
		log.Fatalf("Error geting absolute path for Obsidian directory: %v", err)
	}
	a.ObsidianDir = absObsidianDir
}

func (a *App) Sync() {
	a.SyncAssets()
	a.SyncJournal()
	a.SyncPages()
}

func (a *App) SyncAssets() {
	log.Println("Syncing assets")

}

func (a *App) SyncJournal() {
	log.Println("Syncing journal")

}

func (a *App) SyncPages() {
	log.Println("Syncing pages")

}
