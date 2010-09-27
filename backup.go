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
	LastBackup     uint32
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

func main() {
	conf := new(BackupConf)
	w, e := SetupEnv(conf)

	HandleError("Environment setup", e)
	HandleWarnings("Environment setup", w)

	if e != nil {
		panic("Backup: " + e.String())
	}
	c := make(chan string)
	go TraverseFileTree(conf.Whitelist[0], c)
	for file := range c {
		print(" -> " + file + "\n")
	}
}
