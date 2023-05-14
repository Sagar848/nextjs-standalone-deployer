package vars

type PackageJson struct {
	Name         string                 `json:"name"`
	Version      string                 `json:"version"`
	Private      bool                   `json:"private"`
	Scripts      map[string]interface{} `json:"scripts"`
	Dependencies map[string]interface{} `json:"dependencies"`
}

const NEXT_DEPENDENCY string = "next"
const SCRIPT_START = "start"
const SCRIPT_START_COMMAND = "next start"
const PACKAGE_JSON_FILE = "package.json"
const BUILD_DIR = "build"
const PUBLIC_DIR = "public"
const NEXT_BUILD_DIR = ".next"
const NPM_INSTALL = "npm install"
const YARN_INSTALL = "yarn install"
