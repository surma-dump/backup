package main

import (
	. "../common/_obj/common"
	"flag"
	"fmt"
	"os"
	"runtime"
	"rpc/json"
	"net"
)

func main() {
	defer ErrorHandler()
	path, addr := parseFlags()

	l := net.Listen("tcp", addr)
	for true {
		conn, e := l.Accept()
		if e != nil {
			fmt.Printf("Connection error: %s\n", e.String())
			continue;
		}
		json.ServeConn(conn)
	}
}

type Error struct {
	Description string
	Backtrace bool
}

func printBacktrace() {
	fmt.Printf("Backtrace:\n")
	i := 2
	for _, file, line, ok := runtime.Caller(i); ok; _, file, line, ok = runtime.Caller(i) {
		i++
		fmt.Printf("(%3d) %s:%d\n", i, file, line)
	}
}

func ErrorHandler() {
	if err := recover(); err != nil {
		fmt.Printf("%s\n\n", err.(Error).Description)
		if err.(Error).Backtrace {
			printBacktrace()
		}
		os.Exit(1)
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

	checkFlagValues(path, addr)
	return
}

func checkFlagValues(path, addr string) {
	if IsRegularFile(path) {
		panic(Error{Description: "Path cannot be a file", Backtrace: false})
	}

	if !IsValidAddress(addr) {
		panic(Error{Description: "Invalid listener address", Backtrace: false})
	}
}
