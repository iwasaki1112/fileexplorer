package fileexplorer

import (
	"testing"
)

func TestExplorerFiles(t *testing.T) {
	rootDirectory := "."
	ignoredFilePaths := []string{
		"go.mod",
		"structure.txt",
		".DS_Store",
		".git/",
	}

	ExplorerFiles(rootDirectory, ignoredFilePaths)
}
