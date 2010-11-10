package common

import (
	"io"
	"os"
)

type BackupConfiguration struct {
}

func (b *BackupConfiguration) GetOutputWriter() (io.Writer, os.Error) {
	return os.Open("testfile", os.O_WRONLY | os.O_CREAT, 0755)
}

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
