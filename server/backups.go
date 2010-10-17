package main

import (
	_ "../common/_obj/common"
	"flag"
	"fmt"
	"os"
)

func main() {
	defer errorhandling()
	path, addr := parseFlags()
	_, _ = path, addr
}

type Error struct {
	Description string
}

func errorhandling() {
	if err := recover(); err != nil {
		fmt.Printf("%s\n", err.(Error).Description)
		
	}
}

func parseFlags() (path string, addr string) {
	flag.StringVar(&path, "p", "", "Path to jail the demon to")
	flag.StringVar(&addr, "l", "0.0.0.0:23000", "Address to listen on")
	h := flag.Bool("h", false, "Show Help")

	flag.Parse()

	if *h {
		flag.PrintDefaults()
		panic(Error{Description: ""})
	}

	if path == "" {
		panic(Error{Description: "You have set a path with -p"})
	}

	return
}
