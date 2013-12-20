package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	EMPTY_PATH = ""
)

func replaceBaseName(filePath string, replace string) string {
	return filepath.Join(filepath.Dir(filePath), replace+filepath.Ext(filePath))
}

func baseName(filePath string) string {
	fileName := filepath.Base(filePath)
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func hasExtension(filePath string) bool {
	return len(filepath.Ext(filePath)) > 0
}

func fileExists(filePath string) bool {
	fi, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}

	// if exists and is not a regulare file (i.e. a Dir) returns false
	return fi.Mode().IsRegular()
}

// extract the head of a slice in a safe manner
func safeHead(slice []string) string {
	if len(slice) > 0 {
		return Head(slice)
	}

	return ""
}

// a safe implementation of First that doesn't panic
func safeFirst(f fb, slice []string) string {
	return safeHead(Filter(f, slice))
}

// get the path of the most suitable target for the build
func redoPath(target string) string {
	targets := []string{target + ".do"}
	if hasExtension(target) {
		targets = append(targets, replaceBaseName(target, "default")+".do")
	}
	// we don't use first because First can panic
	return safeFirst(fileExists, targets)
}

// run redo tasks
func redo(target string) {
	// build must be atomic, first we build to a temp file,
	// in case of success we replace the target

	path := redoPath(target)
	if path == EMPTY_PATH {
		log.Fatalf("No .do file found for target '%s'", target)
	}

	log.Print(path)
	tmp := target + "---redoing"
	if out, err := exec.Command("sh", path, "0", baseName(target), tmp).CombinedOutput(); err == nil {
		os.Rename(tmp, target)
	} else {
		log.Fatalf("Redo script returned an error: %s", err)
		// if flags.Verbose -- now it gets never executed
		log.Printf("%s", out)
	}
}

func main() {
	// exec the redo step for each argument
	Each(redo, os.Args[1:])
}
