package internal

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type Config struct {
	FileExtension string
	HomePath      string
	RootPath      string
	FolderPath    string
	EntryPath     string
}

func NewConfig() Config {
	fileExtension := ".md"
	today := time.Now()
	homePath := getHomeDir()
	rootPath := path.Join(homePath, "journal")
	folderPath := path.Join(rootPath, fmt.Sprintf("%d", today.Year()), fmt.Sprintf("%02d", today.Month()))
	entryName := fmt.Sprintf("%02d%s", today.Day(), fileExtension)

	return Config{
		FileExtension: fileExtension,
		HomePath:      homePath,
		RootPath:      rootPath,
		FolderPath:    folderPath,
		EntryPath:     path.Join(folderPath, entryName),
	}
}

func getHomeDir() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get user's home dir: ", err)
	}

	return homePath
}
