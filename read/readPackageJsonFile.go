package read

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/Sagar848/nextjs-standalone-deployer/vars"
)

func ReadPackageJson() (vars.PackageJson, string) {
	var packagejson vars.PackageJson

	packageJsonFile, err := os.Open(vars.PACKAGE_JSON_FILE)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println("Unable to open", vars.PACKAGE_JSON_FILE, "file:", err)
		return vars.PackageJson{}, ""
	}
	log.Println("Successfully Opened", vars.PACKAGE_JSON_FILE)

	byteValue, _ := io.ReadAll(packageJsonFile)
	json.Unmarshal(byteValue, &packagejson)

	deps := packagejson.Dependencies

	var nextVersion string
	if _, ok := deps[vars.NEXT_DEPENDENCY]; ok {
		nextVersion = deps[vars.NEXT_DEPENDENCY].(string)
	} else {
		log.Fatalln(vars.NEXT_DEPENDENCY, "dependency not found!! Terminating")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer packageJsonFile.Close()

	return packagejson, nextVersion
}
