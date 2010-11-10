package main

import (
	. "../common/common"
	"os"
	"json"
	"io"
)

type BackupRPC struct {
	BackupConf *BackupConfiguration
	Output io.Writer
}

func (b *BackupRPC) Set(in *BackupRPCData, out *BackupRPCData) os.Error {
	conf, ok := in.Values["configuration"]
	if !ok {
		return os.NewError("Missing argument \"configuration\"")
	}

	var e os.Error
	b.BackupConf, e = RemarshalBackupConfiguration(conf)
	if e != nil {
		return os.NewError("Expected \"configuration\" to be of type Backup")
	}
	return nil
}

// My BackupRPC struct holds interface{} values. Since their
// actualy dynamic type is lost during transition over the network
// due to the limitations of json, a type assertion to e.g. BackupConfiguration
// is impossible.
// This function tries to put a interface{} into a BackupConfiguration
func RemarshalBackupConfiguration(v interface{}) (*BackupConfiguration, os.Error) {
	bc := new(BackupConfiguration)
	data, e := json.Marshal(v)
	if e != nil {
		return nil, e
	}
	e = json.Unmarshal(data, bc)
	if e != nil {
		return nil, e
	}
	return bc, nil
}
