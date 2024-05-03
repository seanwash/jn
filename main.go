package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

func main() {
	fileExtension := ".md"
	today := time.Now()
	entryName := fmt.Sprintf("%02d%s", today.Day(), fileExtension)
	homePath := getHomeDir()
	rootPath := path.Join(homePath, "journal")
	folderPath := path.Join(rootPath, fmt.Sprintf("%d", today.Year()), fmt.Sprintf("%02d", today.Month()))
	entryPath := path.Join(folderPath, entryName)
	cmd := exec.Command("open", entryPath)
	entryExists, _ := exists(entryPath)

	if entryExists {
		runCmd(cmd)
	} else {
		createEntry(folderPath, entryPath)
		runCmd(cmd)
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

func getHomeDir() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get user's home dir: ", err)
	}

	return homePath
}

func runCmd(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		log.Fatal("Could not run cmd: ", err)
	}
}
