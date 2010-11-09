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
	in.Values["configuration"] = BackupConfiguration{}
	e = c.Call("BackupRPC.Set", in, out)
	if e != nil {
		panic(e.String())
	}
}
