package activity

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/Sagar848/nextjs-standalone-deployer/command"
)

func CopyFolder(sourcePath string, destPath string) {

	var commandToRun string
	switch runtime.GOOS {
	case "windows":
		{
			commandToRun = fmt.Sprint("Copy-Item -Path \"", sourcePath, "\" -Destination \"", destPath, "\" -Recurse -Force")
		}
	default:
		commandToRun = fmt.Sprint("cp ", sourcePath, " ", sourcePath, string(filepath.Separator), destPath)
	}

	_, err := command.ShellCommand(commandToRun)
	if err != nil {
		log.Fatalln("Unable to copy", sourcePath, "to", destPath, "Error:", err)
	} else {
		log.Println("Successfully copied", sourcePath, "to", destPath)
	}
}
