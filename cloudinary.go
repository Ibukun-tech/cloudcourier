package cloudcourier

import (
	"errors"
	"io"

	"github.com/cloudinary/cloudinary-go"
)

type CloudinaryClient struct {
	// Name you give to a file to be stored in cloudinary
	Tag string
	// This is for the transport side responsible for the intercommunication
	Client *cloudinary.Cloudinary
	// This helps to show the cloud provider you want to use
	Provider string
	// The Api key for now I am just using it to access the cloudinary services
	ApiKey string
	// The Api secret which corresponds to the ApiKey
	ApiSecret string
	// This is for cloudinary you need to provide the cloud name for the cloudinary
	CloudName string
}

func newCloudinaryClient(cbb *CloudCourierBridge) (StorageClient, error) {
	if cbb.ApiKey == "" || cbb.ApiSecret == "" || cbb.CloudName == "" {
		return nil, errors.New("incomplete Cloudinary configuration") // TODO: handle properly.
	}
	// To Implement the initialization of cloudinary client
	return &CloudinaryClient{
		ApiKey:    cbb.ApiKey,
		ApiSecret: cbb.ApiSecret,
		CloudName: cbb.CloudName,
	}, nil
}

func (c *CloudinaryClient) UploadFile(filepath string, reader io.Reader) error {
	return nil
}

func (c *CloudinaryClient) DeleteFile(fieldID string) error {
	return nil
}
func (c *CloudinaryClient) ListFiles(directory string) ([]string, error) {
	var result []string
	return result, nil
}
func (c *CloudinaryClient) GetFile(fileID string) (io.Reader, error) {
	var ans io.Reader
	return ans, nil
}
