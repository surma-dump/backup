package main

import ()


type BackupConf struct {
	BackupLocation         string
	NumStacks              uint8
	StackSize              uint8
	CyclicStackArrangement bool
	Whitelist              []string "Folders"
	Blacklist              []string
	Visited                []bool
}

var (
	VERSION = "0.1"
	AUTHORS = "Alexander \"Surma\" Surma <surma@78762.de>"
)

func main() {
	_, _ = SetupEnv()

}
