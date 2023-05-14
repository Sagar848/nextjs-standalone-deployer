package main

import (
	"log"
	"os"

	"github.com/Sagar848/nextjs-standalone-deployer/activity"
	nextPath "github.com/Sagar848/nextjs-standalone-deployer/path"
	"github.com/Sagar848/nextjs-standalone-deployer/read"
	"github.com/Sagar848/nextjs-standalone-deployer/vars"
	"github.com/Sagar848/nextjs-standalone-deployer/write"
)

func main() {
	log.Println("Building next app to build folder started")

	if _, err := os.Stat(vars.BUILD_DIR); os.IsNotExist(err) {
		log.Println("Checking if", vars.BUILD_DIR, "folder exists")
		log.Println(vars.BUILD_DIR, "does not exist, creating new")
		if err := os.Mkdir(vars.BUILD_DIR, os.ModePerm); err != nil {
			log.Fatal(vars.BUILD_DIR, "folder cannot be created", err)
		} else {
			log.Println(vars.BUILD_DIR, "folder created")
		}
	} else {
		log.Println(vars.BUILD_DIR, "exists, hence not creating new")
	}

	packageJson, nextVersion := read.ReadPackageJson()

	log.Println("Details found from package json | projectName =", packageJson.Name, ", projectVersion =", packageJson.Version, ", nextVersion =", nextVersion)

	log.Println("Creating", vars.PACKAGE_JSON_FILE, "in", vars.BUILD_DIR)
	write.CreatePackageJson(packageJson, nextVersion)

	// log.Println("Deleting", vars.PUBLIC_DIR, "directory, before copying new")
	// nextPath.DeleteFromBuild(vars.BUILD_DIR, vars.PUBLIC_DIR)

	log.Println("Copying", vars.PUBLIC_DIR, "to", vars.BUILD_DIR)
	activity.CopyFolder(vars.PUBLIC_DIR, vars.BUILD_DIR)

	// log.Println("Deleting", vars.NEXT_BUILD_DIR, "directory, before copying new")
	// nextPath.DeleteFromBuild(vars.BUILD_DIR, vars.NEXT_BUILD_DIR)

	log.Println("Copying", vars.NEXT_BUILD_DIR, "to", vars.BUILD_DIR)
	activity.CopyFolder(vars.NEXT_BUILD_DIR, vars.BUILD_DIR)

	// nextPath.FindNode()

	log.Println("Running", vars.NPM_INSTALL, "command in", vars.BUILD_DIR)
	nextPath.ChangePathAndRunCommand(vars.BUILD_DIR, vars.NPM_INSTALL)
}
