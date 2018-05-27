package main

import (
	"flag"
	"fmt"
	char "github.com/richardlehane/characterize"
	"os"
	"path/filepath"
)

const version = "fndec-0.0.3"

var (
	file string
	vers bool
)

func init() {
	flag.StringVar(&file, "file", "false", "File to read the filename from.")
	flag.BoolVar(&vers, "version", false, "Return version information.")
}

// Crimped from, https://github.com/richardlehane/siegfried/blob/a3cf617aaec3206f894ac488ae2c5ce1ac2f4230/cmd/sf/sf.go#L82-L101
func ftype(mode os.FileMode) string {
	typ := "unknown"
	switch {
	case mode&os.ModeType == 0:
		typ = "regular file"
	case mode&os.ModeDir == os.ModeDir:
		typ = "directory"
	case mode&os.ModeSymlink == os.ModeSymlink:
		typ = "symlink"
	case mode&os.ModeNamedPipe == os.ModeNamedPipe:
		typ = "named pipe"
	case mode&os.ModeSocket == os.ModeSocket:
		typ = "socket"
	case mode&os.ModeDevice == os.ModeDevice:
		typ = "device"
	}
	return typ
}

func readFileName(path string, fi os.FileInfo, err error) error {
	fmt.Printf("Character encoding for '%s' name: '%s' detected: %s\n",
		ftype(fi.Mode()),
		filepath.Base(path),
		char.Detect([]byte(filepath.Base(path))))
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
