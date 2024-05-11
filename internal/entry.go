package internal

import (
	"fmt"
	"os"
	"os/exec"
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

func (e *Entry) Create() (*os.File, error) {
	err := os.MkdirAll(e.FolderPath, 0700)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(e.FullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (e *Entry) Write(content []byte) error {
	file, err := os.Open(e.FullPath)

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func (e *Entry) Launch() {
	cmd := exec.Command("open", e.FullPath)
	RunCmd(cmd)
}

func (e *Entry) Exists() (bool, error) {
	return Exists(e.FullPath)
}
