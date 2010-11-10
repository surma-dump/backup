package common

import (
)

type BackupConfiguration struct {
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
