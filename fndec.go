package main

import (
	"flag"
	"fmt"
	char "github.com/richardlehane/characterize"
	"os"
	"path/filepath"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "file", "false", "File to read the fname from.")
}

func readFileName(path string, fi os.FileInfo, err error) error {
	fmt.Printf("Character encoding for fname: '%s' detected: %s\n",
		filepath.Base(path), char.Detect([]byte(filepath.Base(path))))
	return nil
}

func main() {
	flag.Parse()
	filepath.Walk(file, readFileName)
}
