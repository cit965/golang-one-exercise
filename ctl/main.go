package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	directory := os.Args[1] // 要统计的目录
	folderCount := 0        // 文件夹数量
	fileCount := 0          // .go文件数量

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if filepath.Ext(path) == ".git" {
				return nil
			}
			if len(info.Name()) == 2 {
				return nil
			}
			if strings.Contains(info.Name(), ".") {
				return nil
			}
			folderCount++
			fmt.Println(info.Name())
		} else if filepath.Ext(path) == ".go" {
			fileCount++
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Folder count:", folderCount)
	fmt.Println(".go file count:", fileCount)
}
