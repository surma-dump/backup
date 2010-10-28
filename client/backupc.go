package main

import (
	. "../common/common"
	"rpc/jsonrpc"
)

func main() {
	c, e := jsonrpc.Dial("tcp", "localhost:23000")
	if e != nil {
		panic(e.String());
	}
	in := NewBackupRPCData()
	out := NewBackupRPCData()
	in.Values["msg"] = "Hallo, Server!\n"
	e = c.Call("BackupRPC.DummyFunc", in, out)
	if e != nil {
		panic(e.String())
	}
	s, ok := out.Values["msg"]
	if ok {
		println(s.(string))
	} else {
		println("aah")
	}
}
