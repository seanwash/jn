package internal

import (
	"fmt"
	"path"
	"time"
)

type Entry struct {
	FileExtension string
	HomePath      string
	RootPath      string
	FolderPath    string
	FullPath      string
}

func NewEntry() Entry {
	fileExtension := ".md"
	today := time.Now()
	homePath := getHomeDir()
	rootPath := path.Join(homePath, "journal")
	folderPath := path.Join(rootPath, fmt.Sprintf("%d", today.Year()), fmt.Sprintf("%02d", today.Month()))
	entryName := fmt.Sprintf("%02d%s", today.Day(), fileExtension)

	return Entry{
		FileExtension: fileExtension,
		HomePath:      homePath,
		RootPath:      rootPath,
		FolderPath:    folderPath,
		FullPath:      path.Join(folderPath, entryName),
	}
}
