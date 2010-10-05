package main

import (
	"fmt"
	"os"
)


type BackupConf struct {
	BackupLocation string
	NumStacks      uint8
	StackSize      uint8
	Whitelist      []string "Folders"
	Blacklist      []string
	Visited        []bool
	IsIncremental  bool
	LastBackup     int64 // nanoseconds since epoch
}

var (
	VERSION = "0.1"
	AUTHORS = "Alexander \"Surma\" Surma <surma@78762.de>"
)

func HandleError(prefix string, e Error) {
	if e != nil {
		ShowHelp(true)
		panic(prefix + ": " + e.String() + "\n")
	}
}

func HandleWarnings(prefix string, w Warnings) {
	if w == nil {
		return
	}

	if len(w) > 0 {
		for _, msg := range w {
			fmt.Fprintf(os.Stderr, "Warning: "+prefix+": "+msg+"\n")
		}
	}
}

func (conf *BackupConf) HasUnvisitedDirectories() bool {
	for _, b := range conf.Visited {
		if !b {
			return true
		}
	}
	return false
}

func (conf *BackupConf) GetUnvisitedDirectory() string {
	for i, b := range conf.Visited {
		if !b {
			return conf.Whitelist[i]
		}
	}
	return ""
}

func (conf *BackupConf) IsBlacklisted(path string) bool {
	blackPrefix := GetLongestPrefix(path, conf.Blacklist)
	whitePrefix := GetLongestPrefix(path, conf.Whitelist)
	return len(blackPrefix) > len(whitePrefix)
}

func (conf *BackupConf) TraverseWhitelist() (<-chan string) {
	out := make(chan string)
	done_signal := make(chan bool)

	// Traverse every path in the Whitelist
	for _, whitepath := range conf.Whitelist {
		go func(path string) {
			allFiles := TraverseFileTree(path)
			for path := range allFiles {
				out <- path
			}
			done_signal <- true
		}(whitepath)
	}

	// Wait far all TraverFileTree() calls do be finished
	// and close channels
	go func() {
		left := len(conf.Whitelist)
		for _ = range done_signal {
			left--
			if  left == 0 {
				close(done_signal)
				close(out)
			}
		}
	}()
	return out
}

func (conf *BackupConf) GetFiles() (out  <-chan *os.File) {
	allFiles := conf.TraverseWhitelist()
	sanitizedFiles := SanitizeFilePaths(allFiles)
	whiteFiles := FilterBlacklistedFiles(sanitizedFiles, func(path string)bool { return conf.IsBlacklisted(path)})
	normalFiles := FilterNormalFiles(whiteFiles)
	uniqueFiles := FilterDuplicates(normalFiles)
	backupFiles := uniqueFiles
	if conf.IsIncremental {
		backupFiles = FilterByTouchDate(uniqueFiles, conf.LastBackup)
	}
	fileHandlers := OpenFiles(backupFiles)
	return fileHandlers
}

func main() {
	conf := new(BackupConf)
	w, e := SetupEnv(conf)

	HandleError("Environment setup", e)
	HandleWarnings("Environment setup", w)

	if e != nil {
		panic("Backup: " + e.String())
	}
	c := conf.GetFiles()
	for file := range c {
		print(" -> " + file.Name() + "\n")
		file.Close()
	}
}
