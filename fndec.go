package main

import (
	"flag"
	"fmt"
	char "github.com/richardlehane/characterize"
	"os"
	"path/filepath"
)

const version = "fndec-0.0.2"

var (
	file string
	vers bool
)

func init() {
	flag.StringVar(&file, "file", "false", "File to read the filename from.")
	flag.BoolVar(&vers, "version", false, "Return version information.")
}

func readFileName(path string, fi os.FileInfo, err error) error {
	fmt.Printf("Character encoding for fname: '%s' detected: %s\n",
		filepath.Base(path), char.Detect([]byte(filepath.Base(path))))
	return nil
}

func main() {
	flag.Parse()
	if vers {
		fmt.Fprintf(os.Stderr, "%s\n", version)
		os.Exit(0)
	} else if flag.NFlag() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: fndec [-file ...]")
		fmt.Fprintln(os.Stderr, "             [Optional: -version]")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Output: Character encoding for fname: {fname} detected: {encoding}")
		fmt.Fprintf(os.Stderr, "Output: [String] '%s ...'\n\n", version)
		flag.Usage()
		os.Exit(0)
	}
	filepath.Walk(file, readFileName)
}
