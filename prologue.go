package main

import (
	"io"
	"flag"
	"os"
	"fmt"
	"strings"
)

func readConfig(path string, conf *BackupConf) (w Warnings, e Error) {
	e = ReadJSONFile(path, conf)
	if e != nil {
		return
	}

	w, e = conf.CheckConfigSanity()
	if e != nil {
		return
	}

	w2, e2 := extrapolateRestData(conf)
	w.Merge(w2)
	if e2 != nil {
		e = e2
	}
	return
}


func extrapolateRestData(conf *BackupConf) (w Warnings, e Error) {
	conf.Visited = make([]bool, len(conf.Whitelist))
	conf.LastBackup, w, e = getBackupDateAndMode(conf)
	return
}


func getBackupDateAndMode(conf *BackupConf) (date int64, w Warnings, e Error) {
	e = conf.CreateBackupTree()
	if e != nil {
		return
	}

	if conf.GetNumBackupsInStack(1) == int(conf.StackSize) {
		conf.ShiftStacks()
	} else {
		date = conf.GetYoungestBackupInStack(1)
	}
	return
}

func ExtractBackupDate(file string) (date int64) {
	n, _ := fmt.Sscanf(file, "incr_%d.tgz", &date)
	if n < 1 {
		fmt.Sscanf(file, "full_%d.tgz", &date)
	}
	return
}

func GetAuthors() []string {
	return strings.Split(AUTHORS, ";", -1)
}

func GetOutput(error bool) io.Writer {
	if error {
		return os.Stderr
	}
	return os.Stdout
}

func ShowHelp(error bool) {
	output := GetOutput(error)

	fmt.Fprintf(output, "backup v%s by:\n", VERSION)
	for _, author := range GetAuthors() {
		fmt.Fprintf(output, "\t%s\n", author)
	}
	flag.PrintDefaults()
}

func SetupEnv(c *BackupConf) (w Warnings, e Error) {
	configFile := flag.String("c", os.Getenv("HOME")+"/.backuprc", "Path to config file")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *help {
		ShowHelp(false)
	}

	if !IsRegularFile(*configFile) {
		e = os.NewError("Config file does not exist or is not a regular file")
		return
	}
	w, e = readConfig(*configFile, c)
	return
}
