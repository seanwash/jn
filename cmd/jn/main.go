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
	templateExists, _ := template.Exists()

	if len(args) > 0 && args[0] == "journal" {
		cmd := exec.Command("open", entry.RootPath)
		internal.RunCmd(cmd)
		return
	}

	if len(args) > 0 && args[0] == "template" {
		if templateExists {
			template.Launch()
			return
		} else {
			_, err := template.Create()
			if err != nil {
				log.Fatal("Could not create template: ", err)
			}
			template.Launch()
			return
		}
	}

	entryExists, err := entry.Exists()
	if err != nil {
		log.Fatal("Could not determine if entry exists: ", err)
	}

	if entryExists {
		entry.Launch()
		return
	}

	_, err = entry.Create()
	if err != nil {
		log.Fatal("Could not create entry: ", err)
	}

	if templateExists {
		templateContent, err := template.Read()

		err = entry.Write(templateContent)
		if err != nil {
			log.Fatal("Could not write template contents to entry: ", err)
		}
	}

	entry.Launch()
}
