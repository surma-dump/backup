package main

import (
	"os"
	paths "path"
	"container/hashmap"
)

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

func FilterNormalFiles(in <-chan string) (<-chan string) {
	out := make(chan string)
	go func() {
		for path := range in {
			file, e := os.Stat(path)
			if e == nil && (file.IsRegular() || file.IsDirectory() || file.IsSymlink()) {
				out <- path
			}
		}
		close(out)
	}()
	return out
}

func inodeHash(v interface{}) int {
	return v.(int)
}

func FilterByInode(in <-chan string) (<-chan string) {
	out := make(chan string)
	go func() {
		inodes := hashmap.New(32, inodeHash)
		for path := range in {
			file, _ := os.Stat(path)
			if !inodes.Containes(file.Ino) {
				inodes.Push(file.Ino)
				out <- path
			}
		}
		close(out)
	}()
	return out
}
