package internal

import (
	"os"
	"os/exec"
	"path"
)

type Template struct {
	FileExtension string
	RootPath      string
	FolderPath    string
	FullPath      string
}

func NewTemplate() Template {
	fileExtension := ".md"
	rootPath := getHomeDir()
	fullPath := path.Join(rootPath, "journal", "template.md")

	return Template{
		FileExtension: fileExtension,
		RootPath:      rootPath,
		FolderPath:    path.Join(rootPath, "journal"),
		FullPath:      fullPath,
	}
}

func (t *Template) Create() (*os.File, error) {
	err := os.MkdirAll(t.FolderPath, 0700)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(t.FullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (t *Template) Read() ([]byte, error) {
	return os.ReadFile(t.FullPath)
}

func (t *Template) Launch() {
	cmd := exec.Command("open", t.FullPath)
	RunCmd(cmd)
}

func (t *Template) Exists() (bool, error) {
	return Exists(t.FullPath)
}
