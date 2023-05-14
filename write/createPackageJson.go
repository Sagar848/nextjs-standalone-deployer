package write

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Sagar848/nextjs-standalone-deployer/vars"
)

func CreatePackageJson(packageJson vars.PackageJson, nextVersion string) {

	packageJsonToBeCreated := packageJson

	dependencyMap := map[string]interface{}{
		vars.NEXT_DEPENDENCY: nextVersion,
	}
	scriptsMap := map[string]interface{}{
		vars.SCRIPT_START: vars.SCRIPT_START_COMMAND,
	}
	packageJsonToBeCreated.Dependencies = dependencyMap
	packageJsonToBeCreated.Scripts = scriptsMap

	file, _ := json.MarshalIndent(packageJsonToBeCreated, "", " ")
	packageJsonPath := fmt.Sprint(vars.BUILD_DIR, string(filepath.Separator), vars.PACKAGE_JSON_FILE)

	err := os.WriteFile(packageJsonPath, file, 0644)
	if err != nil {
		log.Fatalln("Error in writing", vars.PACKAGE_JSON_FILE, "for", vars.BUILD_DIR, "directory")
	} else {
		log.Println(vars.PACKAGE_JSON_FILE, "file created successfully in", vars.BUILD_DIR, "directory")
	}

}
