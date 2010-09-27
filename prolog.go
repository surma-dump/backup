package main

import (
	"io"
	"flag"
	"os"
	"fmt"
	"strings"
	"math"
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

func (conf *BackupConf) CheckConfigSanity() (w Warnings, e Error) {
	for _, path := range conf.Whitelist {
		if !IsNonSpecialFile(path) {
			w.AddWarning("\"" + path + "\" could not be found. It will be ignored")
		}
	}

	return
}

func (conf *BackupConf) GetCurrentStackName() string {
	zeros := fmt.Sprintf("%%0%dd", int(math.Ceil(math.Log10(float64(conf.NumStacks)))))
	return fmt.Sprintf("stack_"+zeros, 0)
}

func extrapolateRestData(conf *BackupConf) (w Warnings, e Error) {
	conf.Visited = make([]bool, len(conf.Whitelist))
	conf.LastBackup, conf.IsIncremental, w, e = getBackupDateAndMode(conf)
	return
}

func getBackupDateAndMode(conf *BackupConf) (date uint32, incr bool, w Warnings, e Error) {
	path := conf.BackupLocation
	if !IsDirectory(path) {
		e = os.Mkdir(path, 0755)
		if e != nil {
			return
		}
	}
	path += "/"+conf.GetCurrentStackName()
	if !IsDirectory(path) {
		e = os.Mkdir(path, 0755)
		if e != nil {
			return
		}
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
