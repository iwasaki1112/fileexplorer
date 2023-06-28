package fileexplorer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ExplorerFiles() {
	ignoredFilePaths := map[string]bool{
		"go.mod":        true,
		"structure.txt": true,
	}

	filePaths, err := listFiles(".", ignoredFilePaths)
	var context string
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}

	for _, filePath := range filePaths {
		content, err := readFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", filePath, err)
		}
		context += "//" + filePath + "\n" + content + "\n"
	}

	saveToFile("./structure.txt", context)
}

func listFiles(rootDirectory string, ignoredFilePaths map[string]bool) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && !ignoredFilePaths[path] {
			filePaths = append(filePaths, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk through directory %s: %w", rootDirectory, err)
	}

	return filePaths, nil
}

func readFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return string(content), nil
}

func saveToFile(filePath string, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}

	return nil
}
