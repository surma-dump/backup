package main

import (
	"json"
	"os"
)

func RegularFileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsRegular()
}

func NonSpecialFileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsRegular() || fileInfo.IsDirectory()
}

func ReadJSONFile(path string, v interface{}) (e Error) {
	file, e1 := os.Open(path, os.O_RDONLY, 0)
	if e1 != nil {
		return e1
	}

	e = json.NewDecoder(file).Decode(v)
	return
}
