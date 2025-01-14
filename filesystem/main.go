package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"slices"
)

type SkipOpts map[string][]string

func main() {
	root := "/Users/muhammadkashif/kashif/courses/golang"
	skipOpts := SkipOpts{
		"dir":  {".git", ".idea"},
		"file": {".gitignore", "api.db"},
	}

	fileSystem := os.DirFS(root)
	err := fs.WalkDir(fileSystem, ".", scanDir(skipOpts))

	if err != nil {
		log.Fatal("DirScan", err)
	}
}

func scanDir(skipOpts SkipOpts) fs.WalkDirFunc {
	var (
		skipDirs  = skipOpts["dir"]
		skipFiles = skipOpts["file"]
	)

	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal("DirWalk", err)
		}

		if d.IsDir() {
			if slices.Contains(skipDirs, path) {
				return fs.SkipDir
			}
		} else if slices.Contains(skipFiles, path) {
			file := formatEntry(d)
			fmt.Println(file)
			return nil
		} else {
			fmt.Println(path)
		}

		return nil
	}
}

func formatEntry(d fs.DirEntry) string {
	file, err := d.Info()
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Name: %s\nSize: %d\nTime: %q\n", file.Name(), file.Size(), file.ModTime())
}
