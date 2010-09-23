package main

import (
	"io"
	"flag"
	"os"
	"fmt"
	"strings"
)

func readConfig(path string, conf *BackupConf) (w Warnings, e Error) {
	e = ReadJSONFile(path, *conf)
	if e != nil {
		return
	}

	w, e = checkConfigSanity(conf)
	return

}

func checkConfigSanity(conf *BackupConf) (w Warnings, e Error) {
	if conf.Visited != nil {
		w.AddWarning("Visited should not be defined by a config file. Its values will be ignored")
	}
	conf.Visited = make([]bool, len(conf.Whitelist))

	for _, path := range conf.Whitelist {
		if !NonSpecialFileExists(path) {
			w.AddWarning("\"" + path + "\" could not be found. It will be ignored")
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

	if !RegularFileExists(*configFile) {
		e = os.NewError("Config file does not exist or is not a regular file")
		return
	}
	return
}
