package activity

import (
	"fmt"
	"log"
	"runtime"

	"github.com/Sagar848/nextjs-standalone-deployer/command"
)

func DeleteFolder(folderPath string) {

	var commandToRun string
	switch runtime.GOOS {
	case "windows":
		{
			commandToRun = fmt.Sprint("Remove-Item -Path \"", folderPath, "\" -Recurse -Force")
		}
	default:
		commandToRun = fmt.Sprint("rm -rf ", folderPath)
	}

	_, err := command.ShellCommand(commandToRun)
	if err != nil {
		log.Fatalln("Unable to delete", folderPath, "Error:", err)
	} else {
		log.Println("Successfully deleted", folderPath)
	}
}
