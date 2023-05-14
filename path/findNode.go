package path

import (
	"fmt"
	"os/exec"
)

func FindNode() {
	node, _ := exec.LookPath("node")
	fmt.Println("node =", node)
}
