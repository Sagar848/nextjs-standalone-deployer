package path

import "os"

func ChangePath(pathToChange string) {

	os.Chdir(pathToChange)
}
