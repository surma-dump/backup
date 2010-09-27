package main

import (
	"json"
	"os"
)

func IsRegularFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsRegular()
}

func IsNonSpecialFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsRegular() || fileInfo.IsDirectory()
}

func IsDirectory(path string) bool {
	d, e := os.Stat(path)
	if e != nil {
		return false
	}

	return d.IsDirectory()
}

func ReadJSONFile(path string, v interface{}) (e Error) {
	file, e1 := os.Open(path, os.O_RDONLY, 0)
	if e1 != nil {
		return e1
	}

	e = json.NewDecoder(file).Decode(v)
	return
}

func GetDirectoryContent(path string) ([]string, Error) {
	f, e := os.Open(path, os.O_RDONLY, 0)
	defer f.Close()
	if e != nil {
		return nil, e
	}
	l, e := f.Readdirnames(-1) // Read all
	return l, e
}

func TraverseFileTree(path string, c chan<- string) {
	TraverseFileTreeRecursive(path, c)
	close(c)
	return
}

func TraverseFileTreeRecursive(path string, c chan<- string) {
	c <- path
	dir := IsDirectory(path)
	if dir {
		f, e := os.Open(path, os.O_RDONLY, 0)
		defer f.Close()
		if e != nil {
			return
		}

		for content, e := f.Readdirnames(1); len(content) > 0 && e == nil; content, e = f.Readdirnames(1) {
			TraverseFileTreeRecursive(path+"/"+content[0], c)
		}
	}
	return
}
