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

func HandleError(prefix string, e Error) {
	if e.Occurred() {
		ShowHelp(true)
		panic(prefix+": "+e.String()+"\n")
	}
}

func main() {
	conf := new(BackupConf)
	_, e := SetupEnv(conf)
	
	HandleError("SetupEnv", e)
	w.DumpWarnings(os.Stderr)

	if e != nil {
		panic("Backup: "+e.String())
	}
}
