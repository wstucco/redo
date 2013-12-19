package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {
	// exec the redo step for each argument
	StringSlice(os.Args[1:]).Map(redo)
}

// run redo tasks
func redo(target string) string {
	// build must be atomic, first we build to a temp file,
	// in case of success we replace the target

	tmp := target + "---redoing"

	if out, err := exec.Command("sh", target+".do", "-", "-", tmp).CombinedOutput(); err == nil {
		os.Rename(tmp, target)
	} else {
		log.Fatalf("Redo script returned an error: %s", err)
		// if flags.Verbose -- now it gets never executed
		log.Printf("%s", out)
	}

	return ""
}

func replaceBaseName(filePath string, replace string) string {
	return path.Join(path.Dir(filePath), replace+path.Ext(filePath))
}
