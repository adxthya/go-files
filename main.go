package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var categories = map[string]string{
	".jpg":  "Images",
	".png":  "Images",
	".jpeg": "Images",
	".mp4":  "Videos",
	".avi":  "Videos",
	".pdf":  "Documents",
	".docx": "Documents",
	".txt":  "Text",
}

func listFiles(directory string) {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(path)
			category := getCategory(ext)
			destFolder := filepath.Join(directory, category)

			moveFiles(path, destFolder)
			fmt.Printf("Moved: %s -> %s\n", path, destFolder)
		}
		return nil
	})
	if err != nil {
		println("err: ", err)
	}
}

func getCategory(ext string) string {
	if folder, exists := categories[ext]; exists {
		return folder
	}
	return "Others"
}

func moveFiles(filePath, destFolder string) {
	err := os.MkdirAll(destFolder, os.ModePerm)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	destPath := filepath.Join(destFolder, filepath.Base(filePath))

	err = os.Rename(filePath, destPath)
	if err != nil {
		fmt.Println("Err in moving file: ", err)
	}
}

func main() {
	listFiles("./testDir")
}
