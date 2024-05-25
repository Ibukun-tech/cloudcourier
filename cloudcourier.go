package cloudcourier

import (
	"errors"
	"io"
)

type ccbCourier string

const (
	Aws        ccbCourier = "aws"
	Gcp        ccbCourier = "gcp"
	Cloudinary ccbCourier = "cloudinary"
)

type CloudCourierBridge struct {
	// used by AWS S3 to specify the Region
	CloudRegion string
	// This helps to show the cloud provider you want to use
	CloudProvider ccbCourier
	// The Api key for now I am just using it to access the cloudinary services
	ApiKey string
	// The Api secret which corresponds to the ApiKey
	ApiSecret string
	// This is for cloudinary you need to provide the cloud name for the cloudinary
	CloudName string
	//We need to specify bucket for other cloud storage providers that make use of it for example s3, Google cloud storage
	CloudBucket string
}

type StorageClient interface {
	UploadFile(filePath string, reader io.Reader) error // Uploads a file from an io.Reader source to the specified path.
	DeleteFile(fileID string) error                     // Deletes a file identified by a unique identifier.
	ListFiles(directory string) ([]string, error)       // Lists files under a specified directory.
	GetFile(fileID string) (io.Reader, error)           // Retrieves a file as an io.Reader by its unique identifier.
}

func NewCloudCourier(ccb *CloudCourierBridge) (StorageClient, error) {
	switch ccb.CloudProvider {
	case Aws:
		return newAWSClient(ccb)
	case Cloudinary:
		return newCloudinaryClient(ccb)
	case Gcp:
		return newGcpClient(ccb)
	default:
		return nil, errors.New("no cloud provider was specified")
	}
}
