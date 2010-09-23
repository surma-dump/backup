package main

import (
	"os"
)

func RegularFileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsRegular()
}
