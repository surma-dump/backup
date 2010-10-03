package main

import (
	"json"
	"os"
	paths "path"
	"strings"
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

func TraverseFileTree(path string) (<-chan string) {
	out := make(chan string)
	go func() {
		TraverseFileTreeRecursive(path, out)
		close(out)
	}()
	return out
}

func TraverseFileTreeRecursive(path string, out chan<- string) {
	out <- path
	dir := IsDirectory(path)
	if dir {
		f, e := os.Open(path, os.O_RDONLY, 0)
		defer f.Close()
		if e != nil {
			return
		}

		for content, e := f.Readdirnames(1); len(content) > 0 && e == nil; content, e = f.Readdirnames(1) {
			TraverseFileTreeRecursive(path+"/"+content[0], out)
		}
	}
	return
}

func FilterBlacklistedFiles(in <-chan string, blackfunc func(string) bool) (<-chan string) {
	out := make(chan string)
	go func() {
		for path := range in {
			if !blackfunc(path) {
				out <- path
			}
		}
		close(out)
	}()
	return out
}

func GetLongestPrefix(s string, prefixes []string) (r string) {
	r = ""
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) && len(prefix) > len(r){
			r = prefix
		}
	}
	return
}

func SanitizeFilePaths(in <-chan string) (<-chan string) {
	out := make(chan string)
	go func() {
		for path := range in {
			out <- paths.Clean(path)
		}
		close(out)
	}()
	return out
}
