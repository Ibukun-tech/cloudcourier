package cloudcourier

import (
	"fmt"
	"io"
)

type cloudCourierProvider int

const (
	AWS cloudCourierProvider = iota
	GCP
)
type Provider interface {
	GetProvider() cloudCourierProvider
}
type serviceFunc func(Provider) (StorageClient, error)

var storeFunc map[cloudCourierProvider]serviceFunc

func RegisterProviderServiceFunc(p cloudCourierProvider, sF serviceFunc) {
	if storeFunc == nil {
		storeFunc = make(map[cloudCourierProvider]serviceFunc)
	}
	storeFunc[p] = sF
}

func NewCloudCourierBridge(c Provider) (StorageClient, error) {
	run, ok := storeFunc[c.GetProvider()]
	if !ok {

		return nil, fmt.Errorf("we do not support this provider")  
	}
	cl, err := run(c)
	if err != nil {
		return nil, fmt.Errorf("failed to create %v client: %v", c.GetProvider(), err)
	}
	return cl, nil
}

type StorageClient interface {
	UploadFile(filePath string, reader io.Reader) error // Uploads a file from an io.Reader source to the specified path.
	DeleteFile(fileID string) error                     // Deletes a file identified by a unique identifier.
	ListFiles(directory string) ([]string, error)       // Lists files under a specified directory.
	GetFile(fileID string) (io.Reader, error)           // Retrieves a file as an io.Reader by its unique identifier.
}

