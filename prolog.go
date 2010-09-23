package main

import (
	"io"
	"flag"
	"os"
	"fmt"
	"strings"
)

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
	for _,author := range GetAuthors() {
		fmt.Fprintf(output, "\t%s\n", author)
	}
}

func SetupEnv() (w Warnings, e Error) {
	configFile := flag.String("c", "~/.backuprc", "Path to config file")
	backupName := flag.String("n", "default", "Name for the template to load")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *help {
		ShowHelp(false)
	}

	if !RegularFileExists(*configFile) {
		e.NewError("Config file does not exist or is not a regular file")
		return
	}
	_ = configFile
	_ = backupName
	return
}
