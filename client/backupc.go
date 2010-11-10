package main

import (
	. "../common/common"
	"rpc/jsonrpc"
)

func main() {
	in := NewBackupRPCData()
	out := NewBackupRPCData()
	in.Values["configuration"] = BackupConfiguration{}
	e = c.Call("BackupRPC.Set", in, out)
	if e != nil {
		panic(e.String())
	}
}
