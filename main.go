package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// gofind: search the file tree starting at the specified directory for a file
// or folder with name matching the given search term.
func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		printUsage()
		log.Fatal("Invalid input")
	}

	// Verify the search directory exists.
	searchDir := args[0]
	_, err := os.Stat(searchDir)
	if os.IsNotExist(err) {
		log.Fatalf("Folder %s does not exist!\n", searchDir)
	}

	absSearchDir := getAbsolutePath(searchDir)

	// Get the search term from command line args.
	searchTerm := args[1]

	// Search!
	fmt.Printf("Searching '%s' for '%s'\n", absSearchDir, searchTerm)
	searchDirectory(absSearchDir, searchTerm)
}

// getAbsolutePath returns the absolute path to the given file or folder.  The
// caller of this function must first verify that the path given relative path
// is valid.
func getAbsolutePath(relativePath string) string {
	var absSearchDir string

	// If path is relative path, prepend the current working directory.
	relativePath = filepath.Clean(relativePath)
	if filepath.IsAbs(relativePath) {
		absSearchDir = relativePath
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("Unable to get current directory\n", err)
		}
		absSearchDir = filepath.Join(cwd, relativePath)
	}

	return absSearchDir
}

// searchDirectory performs a depth-first-search on the given directory searching
// for the given search term, printing the path to the found files.
func searchDirectory(dir string, term string) {
	dirContents, err := ioutil.ReadDir(dir)
	if err != nil {
		// Skip directories we don't have access to.
		return
	}

	for _, child := range dirContents {
		name := child.Name()
		path := filepath.Join(dir, child.Name())

		if strings.Contains(name, term) {
			// It's a match! Print the absolute path to the found file.
			fmt.Println("Found:", path)
		}

		// Recurse on directories.
		if child.IsDir() {
			searchDirectory(path, term)
		}
	}
}

func printUsage() {
	fmt.Println("Usage: gofind <search directory> <search term>")
}
