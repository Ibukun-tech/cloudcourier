package cloudcourier

import (
	file "github.com/Ibukun-tech/cloudcourier/File"
)

type File file.File

type General interface {
	// This is used to upload a single file to the storage provider
	UploadFile(file interface{}) error
	// To check if the client or what will help us in transporting the files are connected
	Connected() bool
	// This is to disconnect incases where there are error
	DisConnect() error
}
