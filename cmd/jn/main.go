package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/seanwash/jn/internal"
)

func main() {
	args := os.Args[1:]
	entry := internal.NewEntry()
	template := internal.NewTemplate()
	templateExists, _ := internal.Exists(template.FullPath)

	if len(args) > 0 && args[0] == "journal" {
		cmd := exec.Command("open", entry.RootPath)
		internal.RunCmd(cmd)
		return
	}

	if len(args) > 0 && args[0] == "template" {
		openTemplateCmd := exec.Command("open", template.FullPath)

		if templateExists {
			internal.RunCmd(openTemplateCmd)
			return
		} else {
			createEntry(template.RootPath, template.FullPath)
			internal.RunCmd(openTemplateCmd)
			return
		}
	}

	openEntryCmd := exec.Command("open", entry.FullPath)
	entryExists, _ := internal.Exists(entry.FullPath)

	if entryExists {
		internal.RunCmd(openEntryCmd)
		return
	}

	if templateExists {
		createEntryFromTemplate(entry.FolderPath, entry.FullPath, template.FullPath)
	} else {
		createEntry(entry.FolderPath, entry.FullPath)
	}

	internal.RunCmd(openEntryCmd)
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

func createEntryFromTemplate(folderPath string, entryPath string, templatePath string) {
	err := os.MkdirAll(folderPath, 0700)
	if err != nil {
		log.Fatal("Could not create new folder", err)
	}

	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatal("Could not read template", err)
	}

	file, err := os.Create(entryPath)
	if err != nil {
		log.Fatal("Could not create new entry", err)
	}

	_, err = file.Write(templateContent)
	if err != nil {
		log.Fatal("Could not write template to new entry", err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal("Could not close new entry", err)
	}
}
