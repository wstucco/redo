package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

type StringSlice []string

func (slice StringSlice) Head() string {
	if len(slice) > 0 {
		return slice[0]
	}

	return ""
}

func (slice StringSlice) Map(fn func(string) string) StringSlice {
	var newSlice StringSlice

	for _, s := range slice {
		newSlice = append(newSlice, fn(s))
	}

	return newSlice
}

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
