package internal

import (
	"log"
	"os"
	"os/exec"
)

func getHomeDir() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get user's home dir: ", err)
	}

	return homePath
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RunCmd(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		log.Fatal("Could not run cmd: ", err)
	}
}
