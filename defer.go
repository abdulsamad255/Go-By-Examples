package main

import (
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "defer.txt")

	f := createFile(path)
	defer closeFile(f) // 👈 scheduled for later

	writeFile(f)
}
