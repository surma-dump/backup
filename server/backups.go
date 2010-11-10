package main

import (
	. "../common/common"
	"flag"
	"fmt"
	"rpc"
	"rpc/jsonrpc"
	"net"
	"git.78762.de/go/error"
)

const (
	DEFAULT_ADDR = "0.0.0.0:23000"
)

func main() {
	defer error.ErrorHandler()
	path, addr := parseFlags()
	_ = path

	l, e := net.Listen("tcp", addr)
	if e != nil {
		panic(e)
	}
	for true {
		conn, e := l.Accept()
		if e != nil {
			fmt.Printf("Connection error: %s\n", e.String())
			continue
		}
		go serveFunctions(conn)
	}
}

func serveFunctions(conn net.Conn) {
	server := rpc.NewServer()
	server.Register(new(BackupRPC))
	server.ServeCodec(jsonrpc.NewServerCodec(conn))
}

func parseFlags() (path string, addr string) {
	flag.StringVar(&path, "p", "", "Path to jail the demon to")
	flag.StringVar(&addr, "l", DEFAULT_ADDR, "Address to listen on")
	h := flag.Bool("h", false, "Show Help")

	flag.Parse()

	if *h {
		flag.PrintDefaults()
		error.Panic("", false)
	}

	checkFlagValues(path, addr)
	return
}

func checkFlagValues(path, addr string) {
	if IsRegularFile(path) {
		error.Panic("File cannot be found", false)
	}

	if !IsValidAddress(addr) {
		error.Panic("Invalid listener address", false)
	}
}
