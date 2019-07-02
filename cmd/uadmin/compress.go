package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/uadmin/uadmin"
)

// Compress takes a folder path and compresses the files inside it
// to a file named publish.tmp
func Compress(path string) {
	outFile, err := os.Create("publish.tmp")
	if err != nil {
		fmt.Println(err)
	}

	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)
	// Add some files to the archive.
	totalCount = countFiles(path)
	currentCount = 0
	addFiles(w, path, "")
	uadmin.Trail(uadmin.OK, "Compressing [%d/%d]", totalCount, totalCount)

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}

var totalCount int
var currentCount int

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	var fullPath string
	for _, file := range files {
		currentCount++
		uadmin.Trail(uadmin.WORKING, "Compressing [%d/%d]", currentCount, totalCount)
		if ignoreFile(file) {
			continue
		}
		fullPath = path.Join(basePath, file.Name())
		//fmt.Println(fullPath)
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(fullPath)
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(path.Join(baseInZip, file.Name()))
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := path.Join(basePath, file.Name())
			//fmt.Println("Recursing and Adding SubDir: " + file.Name())
			//fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, path.Join(baseInZip, file.Name()))
		}
	}
}

// ignoreFile is a list of files that we don't want to publish
func ignoreFile(f os.FileInfo) bool {
	ignoredExt := []string{".db", ".swp", ".salt", ".key", ".uproj"}
	name := f.Name()
	for _, ext := range ignoredExt {
		if strings.HasSuffix(name, ext) {
			return true
		}
	}

	// Exclude Excutable
	if (uint32(f.Mode())&0111) != 0 && !f.IsDir() {
		return true
	}

	return false
}

func countFiles(name string) int {
	files := 0
	root := name

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files++
		return nil
	})
	return files
}
