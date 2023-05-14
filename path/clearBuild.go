package path

import (
	"fmt"
	"path/filepath"

	"github.com/Sagar848/nextjs-standalone-deployer/activity"
)

func DeleteFromBuild(buildPath string, folder string) {
	activity.DeleteFolder(fmt.Sprint(buildPath, string(filepath.Separator), folder))
}
