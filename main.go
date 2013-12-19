package main

import (
	"os/exec"
)

func main() {
	exec.Command("sh", "redo.do").Run()
}
