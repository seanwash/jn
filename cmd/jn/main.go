package main

import (
	"github.com/seanwash/jn/internal"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	entry := internal.NewEntry()

	if len(args) > 0 && args[0] == "journal" {
		cmd := exec.Command("open", entry.RootPath)
		runCmd(cmd)
		return
	}

	openEntryCmd := exec.Command("open", entry.FullPath)
	entryExists, _ := exists(entry.FullPath)

	if entryExists {
		runCmd(openEntryCmd)
	} else {
		createEntry(entry.FolderPath, entry.FullPath)
		runCmd(openEntryCmd)
	}
}

func createEntry(folderPath string, entryPath string) {
	err := os.MkdirAll(folderPath, 0700)
	if err != nil {
		log.Fatal("Could not create new folder", err)
	}

	_, err = os.Create(entryPath)
	if err != nil {
		log.Fatal("Could not create new entry", err)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return true, nil
}

func runCmd(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		log.Fatal("Could not run cmd: ", err)
	}
}
