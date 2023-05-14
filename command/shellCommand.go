package command

import (
	"log"
	"os/exec"
	"runtime"
)

func ShellCommand(command string) ([]byte, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		{
			// cmd = exec.Command("cmd.exe", "/C", "hostname")
			ps, _ := exec.LookPath("powershell.exe")
			// fmt.Println("ps =", ps)

			cmd = exec.Command(ps, "-NoProfile", "-NonInteractive", command)
		}
	default:
		cmd = exec.Command("sh", "-c", command)
	}

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error in getting output for command", command, ", Error is: ", err)
	}
	// fmt.Printf("%s\n", stdoutStderr)
	return stdoutStderr, err
}
