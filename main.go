package main

import (
	"log"
	"os"
	"os/exec"
)

type StringSlice []string

func (slice StringSlice) Each(fn func(string)) {
	for _, s := range slice {
		fn(s)
	}
}

func main() {
	// exec the redo step for each argument
	StringSlice(os.Args[1:]).Each(redo)
}

// run redo tasks
func redo(target string) {
	// build must be atomic, first we build to a temp file,
	// in case of success we replace the target

	tmp := target + "---redoing"

	if out, err := exec.Command("sh", target+".do", "-", "-", tmp).CombinedOutput(); err == nil {
		os.Rename(tmp, target)
	} else {
		log.Fatalf("Redo script returned an error: %s", err)
		// if flags.Verbose
		log.Printf("%s", out)
	}
}
