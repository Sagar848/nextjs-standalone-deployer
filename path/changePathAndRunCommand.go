package path

import (
	"log"

	cmd "github.com/Sagar848/nextjs-standalone-deployer/command"
)

func ChangePathAndRunCommand(pathToChange string, command string) {

	ChangePath(pathToChange)

	_, err := cmd.ConsoleCommand(command)
	if err != nil {
		log.Fatalln("Unable to run", command, "in", pathToChange)
	} else {
		log.Println(command, "executed successfully")
	}
}
