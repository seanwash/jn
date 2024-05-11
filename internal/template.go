package internal

import "path"

type Template struct {
	FileExtension string
	RootPath      string
	FullPath      string
}

func NewTemplate() Template {
	fileExtension := ".md"
	rootPath := getHomeDir()
	fullPath := path.Join(rootPath, "journal", "template.md")

	return Template{
		FileExtension: fileExtension,
		RootPath:      rootPath,
		FullPath:      fullPath,
	}
}
