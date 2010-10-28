package common

import (
	"os"
)

type Backup struct {
}

type BackupRPC Backup

type BackupRPCData struct {
	Values map[string]interface{}
}

func NewBackupRPCData() (r *BackupRPCData) {
	r = new(BackupRPCData)
	r.Values = make(map[string]interface{})
	return
}

func InitBackupRPCData() (r BackupRPCData) {
	return *NewBackupRPCData()
}

func (b *BackupRPC) DummyFunc(in *BackupRPCData, out *BackupRPCData) os.Error {
	s, ok := in.Values["msg"]
	if !ok {
		return os.NewError("Missing argument \"msg\"")
	}
	println(s.(string))
	*out = InitBackupRPCData()
	out.Values["msg"] = "Hallo, Client!\n"
	return nil
}
